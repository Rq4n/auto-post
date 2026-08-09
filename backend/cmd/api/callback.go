package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

func beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}

func getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(w, r)
		return
	}

	session, err := gothic.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return
	}

	session.Values["user_id"] = user.UserID

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return

	}

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
