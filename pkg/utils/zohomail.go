package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/cave/configs"
)

type Mail struct {
	fromAddress  string `json:"fromAddress"`
	toAddress    string `json:"toAddress"`
	subject      string `json:"subject"`
	content      string `json:"content"`
	askReceipt   string `json:"askReceipt"`
	code         string `json:"code"`
	
	Response     []byte
	accessToken  string
	refreshToken string
	accountID    string
	template     string
	config       configs.Mail
}

func NewMailer(config *configs.AppConfig) *Mail {
	return &Mail{
		config: config.Mail,
	}
}

func (m *Mail) Send() error {
	apiUrl := fmt.Sprintf("https://mail.zoho.com/api/accounts/%s/messages", m.accountID)
	resource := "/user/"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"

	err := m.parseTemplate()
	if err != nil {
		fmt.Println("mail: Error parsing template: ", err)
		return err
	}

	body, err := json.Marshal(m)
	if err != nil {
		fmt.Println("mail: Error reading mail body")
		return err
	}

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(body)) // URL-encoded payload
	r.Header.Add("Accept", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Zoho-oauthtoken %s", os.Getenv("ZOHOMAIL_TOKEN")))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("mail: Error sending mail")
		return err
	}

	defer resp.Body.Close()
	m.Response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("mail: Errored reading response body", err)
		return err
	}
	return nil
}

func (m *Mail) parseTemplate() error {
	t, err := template.ParseFiles(m.template)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, m); err != nil {
		return err
	}
	m.content = buf.String()
	return nil

}

func (m *Mail) getAuthCode() error {
	apiUrl := "https://accounts.zoho.com/oauth/v2/"
	resource := "/auth"
	data := url.Values{}

	data.Set("client_id", m.config.ClientID)
	data.Set("client_secret", m.config.ClientSecret)
	data.Set("response_type", "code")
	data.Set("redirect_uri", m.config.ClientSecret)
	data.Set("scope", m.config.Scope)
	data.Set("access_type", "offline")
	data.Set("state", "14142791")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("mail: Error sending mail")
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("mail: Errored reading response body")
		return err
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("mail: Errored unmashalling response body")
		return err
	}
	return nil
}

func (m *Mail) getAccessToken() error {
	apiUrl := "https://accounts.zoho.com/oauth/v2/"
	resource := "/auth"
	data := url.Values{}

	data.Set("client_id", m.config.ClientID)
	data.Set("client_secret", m.config.ClientSecret)
	data.Set("response_type", "code")
	data.Set("redirect_uri", m.config.ClientSecret)
	data.Set("scope", m.config.Scope)
	data.Set("access_type", m.config.ClientSecret)
	data.Set("state", "14142791")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "https://api.com/user/"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("mail: Error sending mail")
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("mail: Errored reading response body")
		return err
	}

	err = json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println("mail: Errored unmashalling response body")
		return err
	}
	return nil

}
