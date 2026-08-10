package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func beginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}

func (h *UserHandler) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	googleUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(w, r)
		return
	}
	user, err := h.userService.GetUserByGoogleID(r.Context(), googleUser.Email)
	if err != nil {
		h.userService.CreateUser(
			r.Context(),
			googleUser.Email,
			googleUser.UserID,
		)
		if err != nil {
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
	}

	session, err := gothic.Store.Get(r, "session")
	if err != nil {
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return
	}

	session.Values["user_id"] = user.ID

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return

	}

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
