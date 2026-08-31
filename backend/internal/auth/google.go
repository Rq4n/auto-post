// Package auth
package auth

import (
	"net/http"
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
	if isProd {
		store.Options.SameSite = http.SameSiteNoneMode
	} else {
		store.Options.SameSite = http.SameSiteLaxMode
	}

	gothic.Store = store

	goth.UseProviders(
		twitterv2.New(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWITTER_CLIENT_SECRET"), "http://127.0.0.1:8080/social/connections/twitter/callback"),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
}
