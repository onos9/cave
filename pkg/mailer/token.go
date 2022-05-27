package mailer

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cave/config"
	"github.com/cave/pkg/database"
)

func (m *Mail) getNewToken() (string, error) {
	cfg := config.GetMailConfig()
	rdb := database.RedisClient(0)
	defer rdb.Close()

	at, err := rdb.Get(ctx, "accessToken").Result()
	if err == nil && at != "" {
		m.AccessToken = at
		return at, nil
	}

	rt, err := rdb.Get(ctx, "zohoRefreshToken").Result()
	if err != nil {
		return "", errors.New("zohoRefreshToken: " + err.Error())
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

func (m *Mail) RequestTokens(code string) (string, error) {
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
