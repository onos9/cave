package mailer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/cave/config"
	"github.com/cave/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type Mail struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ApiDomain    string `json:"api_domain"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Error        string `json:"error"`
}

var ctx = context.Background()

func (m *Mail) SendMail(mail fiber.Map) (fiber.Map, error) {
	cfg := config.GetMailConfig()
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s/messages", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	tpl, err := utils.ParseTemplate(mail["content"].(fiber.Map))
	if err != nil {
		return fiber.Map{}, err
	}

	log.Println(tpl)

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

func (m *Mail) GetCredential() (fiber.Map, error) {
	cfg := config.GetMailConfig()
	apiUrl := "https://mail.zoho.com"
	resource := fmt.Sprintf("api/accounts/%s", cfg.AccountID)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	_, err := m.getNewToken()
	if err != nil {
		return fiber.Map{}, err
	}

	r, _ := http.NewRequest(http.MethodGet, urlStr, nil)
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", m.AccessToken))

	var resp fiber.Map
	err = doRequest(r, &resp)
	if err != nil {
		return fiber.Map{}, err
	}

	return resp, nil
}

func (m *Mail) GetMailConfig() fiber.Map {
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
