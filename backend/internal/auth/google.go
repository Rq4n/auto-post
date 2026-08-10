// Package auth
package auth

import (
	"os"

	"github.com/Rq4n/autopost/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitterv2"
)

var Store *sessions.Store

const (
	key    = "0123456789abcdef0123456789abcdef"
	maxAge = 86400 * 30 // 30 days
	isProd = false      // Set to true when serving over https

)

func NewAuth(cfg config.GoogleOAuth, twitter config.TwitterOAuth) {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		twitterv2.New(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"), "http://localhost:8080/auth/twitterv2/callback"),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/google/callback"),
	)

	// goth.UseProviders(
	// 	google.New(
	// 		cfg.ClientID,
	// 		cfg.ClientSecret,
	// 		cfg.RedirectURL,
	// 		cfg.Scopes...,
	// 	),
	//
	// 	twitterv2.New(
	// 		twitter.ClientID,
	// 		twitter.ClientSecret,
	// 		"http://localhost:8080/auth/twitterv2/callback",
	// 	),
	// )
}
