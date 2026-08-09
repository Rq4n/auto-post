// Package auth
package auth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitterv2"
)

const (
	key    = "randomString"
	maxAge = 86400 * 30 // 30 days
	isProd = false      // Set to true when serving over https
)

func NewAuth() {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		twitterv2.New(os.Getenv("TWITTER_CLIENT_ID"), os.Getenv("TWIITER_CLIENT_SECRET"), "http://localhost:3000/auth/twitterv2/callback"),
		google.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:3000/auth/google/callback"),
	)
}
