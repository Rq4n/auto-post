// Package handler
package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth/gothic"
)

type TwitterHandler struct {
	socialService service.SocialService
}

func NewTwitterHandler(socialService service.SocialService) *TwitterHandler {
	return &TwitterHandler{
		socialService: socialService,
	}
}

func mapToGothProvider(provider string) (string, error) {
	switch provider {
	case "twitter":
		return "twitterv2", nil
	default:
		return "", fmt.Errorf("unsupported provider: %s", provider)
	}
}

func (t *TwitterHandler) BeginProviderAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	gothProvider := "twitterv2"

	switch provider {
	case "twitter":
		gothProvider = "twitterv2"
	default:
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}

	r = r.WithContext(
		context.WithValue(r.Context(), "provider", gothProvider),
	)

	gothic.BeginAuthHandler(w, r)
}

func (t *TwitterHandler) HandleTwitterCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	gothProvider, err := mapToGothProvider(provider)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Informa ao Goth qual provider está sendo finalizado.
	ctx := context.WithValue(r.Context(), "provider", gothProvider)

	r = r.WithContext(ctx)

	// Finaliza o OAuth.
	// Aqui o Goth retorna os dados da conta social.
	pvUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("failed to complete %s oauth: %v", gothProvider, err)
		http.Error(w, "failed to complete oauth", http.StatusInternalServerError)
		return
	}

	// Pega o usuário do AutoPost que está logado.
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		log.Printf("failed to get user_id from session context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Converte o userID do AutoPost para UUID.
	parseID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	sID := pgtype.UUID{
		Bytes: parseID,
		Valid: true,
	}

	// sID              = usuario Autopost
	// provider         = twitter/linkedin/bluesky
	// pvUser.UserID    = usuário da rede social
	// pvUser.AccessToken
	// pvUser.RefreshToken
	// pvUser.ExpiresAt
	//
	// Service salva tudo
	_, err = t.socialService.ConnectNewProvider(r.Context(), sID, provider, pvUser)
	if err != nil {
		log.Printf("failed to create social connection: %v", err)
		http.Error(w, "failed to save social connection", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
}
