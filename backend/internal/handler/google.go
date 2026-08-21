package handler

import (
	"context"
	"log"
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

func BeginAuthProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	gothic.BeginAuthHandler(w, r)
}

func (h *UserHandler) GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)

	googleUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.userService.GetUserByGoogleID(
		r.Context(),
		googleUser.Email,
	)
	if err != nil {
		user, err = h.userService.CreateUser(
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
		log.Printf("failed to get session %v", err)
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return
	}

	session.Values["user_id"] = user.ID.String()

	err = session.Save(r, w)
	if err != nil {
		log.Printf("failed to save session %v", err)
		http.Error(w, "sessions error", http.StatusInternalServerError)
		return
	}

	log.Printf("user created %v", googleUser.UserID)
	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
