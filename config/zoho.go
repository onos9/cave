package config

import "os"

// RedisConfig object
type ZohoMail struct {
	FromEmail       string `env:"FROM_EMAIL"`
	ClientID        string `env:"CLIENT_ID"`
	ClientSecret    string `env:"CLIENT_SECRET"`
	ResponseType    string `env:"RESPONSE_TYPE"`
	RedirectURL     string `env:"REDIRECT_URL"`
	AccountID       string `env:"ACCOUNT_ID"`
	Scope           string `env:"SCOPE"`
	MailLaizyKey    string `env:"MAIL_LAIZY_KEY"`
	MailLaizySecret string `env:"MAIL_LAIZY_SECRET"`
	ZOID            string `env:"ZOID"`
	State           string `env:"STATE"`
}

// GetRedisConfig returns RedisConfig object
func GetMailConfig() ZohoMail {
	return ZohoMail{
		FromEmail:       os.Getenv("EMAIL_FROM"),
		ClientID:        os.Getenv("CLIENT_ID"),
		ClientSecret:    os.Getenv("CLIENT_SECRET"),
		RedirectURL:     os.Getenv("REDIRECT_URL"),
		AccountID:       os.Getenv("ACCOUNT_ID"),
		MailLaizyKey:    os.Getenv("MAIL_LAIZY_KEY"),
		MailLaizySecret: os.Getenv("MAIL_LAIZY_SECRETE"),
		Scope:           os.Getenv("SCOPE"),
		ZOID:            os.Getenv("ZOID"),
		State:           os.Getenv("STATE"),
	}
}
