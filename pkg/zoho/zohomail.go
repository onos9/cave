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
	"github.com/gofiber/fiber/v2"
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

func parseTemplate(m fiber.Map) (string, error) {
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

func (m *Mailer) getNewToken() (string, error) {
	cfg := config.GetMailConfig()
	rdb := database.RedisClient(0)
	defer rdb.Close()

	at, err := rdb.Get(ctx, "accessToken").Result()
	if err == nil || at != "" {
		m.AccessToken = at
		return at, nil
	}

	rt, err := rdb.Get(ctx, "zohoRefreshToken").Result()
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

	err = doRequest(r, &m)
	if err != nil || m.Error != "" {
		return m.Error, err
	}

	expireTime := time.Duration(m.ExpiresIn) * time.Second
	err = rdb.Set(ctx, "accessToken", m.AccessToken, expireTime).Err()
	if err != nil {
		return "", err
	}

	return m.AccessToken, nil
}

func (m *Mailer) RequestTokens(code string) (string, error) {
	cfg := config.GetMailConfig()
	
	rdb := database.RedisClient(0)
	defer rdb.Close()

	err := rdb.Set(ctx, "zoho_code", code, 0).Err()
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

	err = doRequest(r, &m)
	if err != nil {
		return "", err
	}

	err = rdb.Set(ctx, "zohoRefreshToken", m.RefreshToken, 0).Err()
	if err != nil {
		return "", err
	}

	expireTime := time.Duration(m.ExpiresIn) * time.Second
	err = rdb.Set(ctx, "accessToken", m.AccessToken, expireTime).Err()
	if err != nil {
		return "", err
	}

	return m.AccessToken, nil
}

func (m *Mailer) GetCredential() (fiber.Map, error) {
	cfg := config.GetMailConfig()
	apiUrl := "https://m.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", m.AccessToken))

	var resp fiber.Map
	err := doRequest(r, &resp)
	if err != nil {
		return fiber.Map{}, err
	}

	return resp, nil
}

func (m *Mailer) SendMail(mail fiber.Map) (fiber.Map, error) {
	cfg := config.GetMailConfig()
	apiUrl := "https://m.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s/messages", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	tpl, err := parseTemplate(mail)
	if err != nil {
		return fiber.Map{}, err
	}

	mail["content"] = tpl

	json_data, err := json.Marshal(mail)
	if err != nil {
		return fiber.Map{}, err
	}

	_, err = m.getNewToken()
	if err != nil {
		return fiber.Map{}, err
	}

	r, _ := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(json_data))
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", m.AccessToken))
	r.Header.Add("Content-Type", "application/json")

	var resp fiber.Map
	err = doRequest(r, &resp)
	if err != nil {
		return fiber.Map{}, err
	}

	return resp, nil
}

func (m *Mailer) GetMailConfig() fiber.Map {
	cfg := config.GetMailConfig()
	return fiber.Map{
		"client_id":     cfg.ClientID,
		"redirect_uri":  cfg.RedirectURL,
		"scope":         cfg.Scope,
		"response_type": "code",
		"access_type":   "offline",
		"prompt":        "consent",
	}
}

