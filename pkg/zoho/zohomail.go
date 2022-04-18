package zoho

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/cave/config"
	"github.com/cave/pkg/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var (
	mail *Mailer
	cfg  config.ZohoMail
	RdDB *redis.Client
)

type Mailer struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ApiDomain    string `json:"api_domain"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`
}

var ctx = context.Background()

func NewMailer(conf config.Config) {
	cfg = conf.ZohoMail
	mail = &Mailer{}
	RdDB = database.RdDB
}

func doRequest(r *http.Request, v interface{}) error {
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}
	return nil
}

func RequestTokens(code string) (string, error) {
	err := RdDB.Set(ctx, "zoho_code", code, 0).Err()
	if err != nil {
		return "", err
	}

	apiUrl := "https://accounts.zoho.com"
	resource := "/oauth/v2/token"
	data := url.Values{}

	data.Set("code", code)
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("redirect_uri", cfg.RedirectURL)
	data.Set("grant_type", "authorization_code")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	err = doRequest(r, &mail)
	if err != nil {
		return "", err
	}

	err = RdDB.Set(ctx, "zohoRefreshToken", mail.RefreshToken, 0).Err()
	if err != nil {
		return "", err
	}

	expireTime := time.Duration(mail.ExpiresIn) * time.Second
	err = RdDB.Set(ctx, "accessToken", mail.AccessToken, expireTime).Err()
	if err != nil {
		return "", err
	}

	return mail.AccessToken, nil
}

func GetNewToken() (string, error) {
	at, err := RdDB.Get(ctx, "accessToken").Result()
	if err == nil {
		mail.AccessToken = at
		return at, nil
	}

	rt, err := RdDB.Get(ctx, "zohoRefreshToken").Result()
	if err != nil {
		return "", err
	}

	apiUrl := "https://accounts.zoho.com"
	resource := "/oauth/v2/token"
	data := url.Values{}

	data.Set("refresh_token", rt)
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("redirect_uri", cfg.RedirectURL)
	data.Set("grant_type", "refresh_token")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	err = doRequest(r, &mail)
	if err != nil {
		return "", err
	}

	expireTime := time.Duration(mail.ExpiresIn) * time.Second
	err = RdDB.Set(ctx, "accessToken", mail.AccessToken, expireTime).Err()
	if err != nil {
		return "", err
	}

	return mail.AccessToken, nil
}

func GetCredential() (fiber.Map, error) {
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", mail.AccessToken))

	var resp fiber.Map
	err := doRequest(r, &resp)
	if err != nil {
		return fiber.Map{}, err
	}

	return resp, nil
}

func SendMail(m fiber.Map) (fiber.Map, error) {
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s/messages", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	tpl, err := ParseTemplate(m)
	if err != nil {
		return fiber.Map{}, err
	}

	m["content"] = tpl

	json_data, err := json.Marshal(m)
	if err != nil {
		return fiber.Map{}, err
	}

	_, err = GetNewToken()
	if err != nil {
		return fiber.Map{}, err
	}

	r, _ := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(json_data))
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", mail.AccessToken))
	r.Header.Add("Content-Type", "application/json")

	var resp fiber.Map
	err = doRequest(r, &resp)
	if err != nil {
		return fiber.Map{}, err
	}

	return resp, nil
}

func ParseTemplate(m fiber.Map) (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("can not get filename")
	}

	dir := path.Dir(filename)
	filePath := fmt.Sprintf("%s/templates/signup.html", dir)
	t, err := template.ParseFiles(filePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, m); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetZohoMailConfig() fiber.Map {
	return fiber.Map{
		"client_id":     cfg.ClientID,
		"redirect_uri":  cfg.RedirectURL,
		"scope":         cfg.Scope,
		"response_type": "code",
		"access_type":   "offline",
		"prompt":        "consent",
	}
}
