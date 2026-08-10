// Package config
package config

import (
	"os"

	"github.com/Rq4n/autopost/pkg/env"
)

type Config struct {
	Port      string
	DB        PostgresConfig
	Providers Providers
}

type Providers struct {
	Google  GoogleOAuth
	Twitter TwitterOAuth
}

type GoogleOAuth struct {
	ClientSecret string
	ClientID     string
	RedirectURL  string
	// Scopes       []string
}

type TwitterOAuth struct {
	ClientSecret string
	ClientID     string
	BearerToken  string
}

type PostgresConfig struct {
	Username    string
	Password    string
	DB          string
	Host        string
	Port        string
	MaxIdleTime string
	MinIdleConn int
	MaxOpenConn int
	SSLMode     string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		Providers: Providers{
			Google: GoogleOAuth{
				ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
				ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
				RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
				// Scopes: []string{
				// 	"openid",
				// 	"email",
				// 	"profile",
				// },
			},
			Twitter: TwitterOAuth{
				ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
				ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
				BearerToken:  os.Getenv("TWITTER_BEARER_TOKEN"),
			},
		},
		DB: PostgresConfig{
			Username:    os.Getenv("POSTGRES_USER"),
			Password:    os.Getenv("POSTGRES_PASSWORD"),
			Host:        os.Getenv("POSTGRES_HOST"),
			Port:        os.Getenv("POSTGRES_PORT"),
			DB:          os.Getenv("POSTGRES_DB"),
			MaxIdleTime: os.Getenv("POSTGRES_MAX_IDLE_TIME"),
			MaxOpenConn: env.GetInt("POSTGRES_MAX_OPEN_CONN", 0),
			MinIdleConn: env.GetInt("POSTGRES_MIN_IDLE_CONN", 0),
			SSLMode:     os.Getenv("POSTGRES_SSLMODE"),
		},
	}

	return cfg, nil
}
