package zoho

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	configs "github.com/cave/config"
	"github.com/cave/pkg/database"
	"github.com/go-redis/redis/v8"
	"github.com/mailazy/mailazy-go"
)

var (
	mail *Mailer
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Mailer struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
	AskReceip   string `json:"askReceip"`

	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ApiDomain    string `json:"api_domain"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`

	body    string
	cfg     configs.Config
	redis   *redis.Client
	Respons map[string]interface{}
}

var ctx = context.Background()

func NewMailer(config configs.Config) {
	mail = &Mailer{
		cfg:   config,
		redis: database.RdDB,
	}
}

func ParseTemplate() error {
	t, err := template.ParseFiles("template.html")
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, mail); err != nil {
		return err
	}
	mail.body = buf.String()
	return nil

}

func RequestTokens(code string) error {

	// we can call set with a `Key` and a `Value`.
	err := mail.redis.Set(ctx, "zoho_code", code, 0).Err()
	if err != nil {
		return err
	}

	apiUrl := "https://accounts.zoho.com"
	resource := "/oauth/v2/token"
	data := url.Values{}

	mCfg := mail.cfg.ZohoMail
	data.Set("code", code)
	data.Set("client_id", mCfg.ClientID)
	data.Set("client_secret", mCfg.ClientSecret)
	data.Set("redirect_uri", mCfg.RedirectURL)
	data.Set("grant_type", "authorization_code")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	// r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mail)
	if err != nil {
		return err
	}

	err = mail.redis.Set(ctx, "zohoRefreshToken", mail.RefreshToken, 0).Err()
	if err != nil {
		return err
	}

	mail.Respons = map[string]interface{}{
		"success": true,
	}
	return nil
}

func GetNewToken() error {

	// we can call set with a `Key` and a `Value`.
	rt, err := mail.redis.Get(ctx, "zohoRefreshToken").Result()
	if err != nil {
		return err
	}

	apiUrl := "https://accounts.zoho.com"
	resource := "/oauth/v2/token"
	data := url.Values{}

	mCfg := mail.cfg.ZohoMail
	data.Set("refresh_token", rt)
	data.Set("client_id", mCfg.ClientID)
	data.Set("client_secret", mCfg.ClientSecret)
	data.Set("redirect_uri", mCfg.RedirectURL)
	data.Set("grant_type", "refresh_token")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mail)
	if err != nil {
		return err
	}

	mail.Respons = map[string]interface{}{
		"access_token": mail.AccessToken,
	}

	return nil
}

func GetCredential() error {
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s", mail.cfg.ZohoMail.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", mail.AccessToken))

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mail.Respons)
	if err != nil {
		return err
	}

	return nil

}

func ZohoSend(m Mailer) error {
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s/messages", mail.cfg.ZohoMail.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	mailData := map[string]string{
		"fromAddress": mail.cfg.ZohoMail.FromEmail,
		"toAddress":   m.ToAddress,
		"subject":     m.Subject,
		"content":     m.Content,
		"askReceip":   m.AskReceip,
	}

	json_data, err := json.Marshal(mailData)
	if err != nil {
		return err
	}
	log.Println(string(json_data))

	//m.GetNewToken()

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(json_data))
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", mail.AccessToken))
	r.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &mail.Respons)
	if err != nil {
		return err
	}

	return nil

}

func Send(m Mailer) *mailazy.SendMailError {
	// err := parseTemplate()
	// if err != nil {
	// 	return err
	// }
	mCfg := mail.cfg
	senderClient := mailazy.NewSenderClient(mCfg.MailLaizyKey, mCfg.MailLaizySecret)
	to := m.ToAddress
	from := "Adullam <no-reply@adullam.ng>"
	subject := m.Subject
	textContent := m.body
	htmlContent := ""
	req := mailazy.NewSendMailRequest(to, from, subject, textContent, htmlContent)

	resp, err := senderClient.Send(req)
	if err != nil {
		return err
	}

	mail.Respons = map[string]interface{}{
		"mailRespons": resp.Message,
	}

	return nil
}
