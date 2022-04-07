package configs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const (
	prod = "production"
)

// Config object
type Config struct {
	Env              string        `env:"ENV"`
	Mongo            MongoDBConfig `json:"mongo"`
	Redis            RedisConfig   `json:"redis"`
	ZohoMail         ZohoMail      `json:"zohomail"`
	JWTAccessSecret  string        `env:"JWT_ACCESS_SIGN_KEY"`
	JWTRefreshSecret string        `env:"JWT_REFRESH_SIGN_KEY"`
	JWTIssuer        string        `env:"JWT_ISSUER"`
	Host             string        `env:"APP_HOST"`
	Port             string        `env:"APP_PORT"`
	MailLaizySecret  string        `env:"MAIL_LAIZY_SECRETE"`
	MailLaizyKey     string        `env:"MAIL_LAIZY_KEY"`
}

// IsProd Checks if env is production
func (c Config) IsProd() bool {
	return c.Env == prod
}

// LoadConfig gets config from .env
func LoadConfig() {
	currentPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	environmentPath := filepath.Join(currentPath, ".env")

	if err := godotenv.Load(environmentPath); err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

// GetConfig gets all config for the application
func GetConfig() Config {
	return Config{
		Env:      os.Getenv("ENV"),
		Mongo:    GetMongoDBConfig(),
		Redis:    GetRedisConfig(),
		ZohoMail: GetZohoMailConfig(),

		JWTAccessSecret:  os.Getenv("JWT_ACCESS_SIGN_KEY"),
		JWTRefreshSecret: os.Getenv("JWT_REFRESH_SIGN_KEY"),
		JWTIssuer:        os.Getenv("JWT_ISSUER"),
		Host:             os.Getenv("APP_HOST"),
		Port:             os.Getenv("APP_PORT"),
		MailLaizySecret:  os.Getenv("MAIL_LAIZY_SECRETE"),
		MailLaizyKey:     os.Getenv("MAIL_LAIZY_KEY"),
	}
}
