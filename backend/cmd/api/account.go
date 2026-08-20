package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth/gothic"
)

type SocialHandler struct {
	socialService service.SocialService
}

func NewSocialHandler(socialService service.SocialService) *SocialHandler {
	return &SocialHandler{
		socialService: socialService,
	}
}

type ProviderPayload struct {
	Provider string
}

func (s *SocialHandler) BeginProviderAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	gothic.BeginAuthHandler(w, r)
}

func (s *SocialHandler) handleConnectNewProvider(w http.ResponseWriter, r *http.Request) {
	var req ProviderPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var gothProvider string

	switch req.Provider {
	case "twitter":
		gothProvider = "twitterv2"

	case "linkedin":
		gothProvider = "linkedin"

	case "bluesky":
		gothProvider = "bluesky"

	default:
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "provider", gothProvider))

	gothic.BeginAuthHandler(w, r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (s *SocialHandler) handleProviderCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Converte o provider da URL para o nome usado pelo Goth.
	var gothProvider string

	switch provider {
	case "twitter":
		gothProvider = "twitterv2"

	case "linkedin":
		gothProvider = "linkedin"

	case "bluesky":
		gothProvider = "bluesky"

	default:
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}

	// Informa ao Goth qual provider está sendo finalizado.
	ctx := context.WithValue(r.Context(), "provider", gothProvider)

	r = r.WithContext(ctx)

	// Finaliza o OAuth.
	// Aqui o Goth retorna os dados da conta social.
	pvUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("failed to complete %s oauth: %v", provider, err)
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
	_, err = s.socialService.ConnectNewProvider( r.Context(), sID, provider, pvUser)
	if err != nil {
		log.Printf("failed to create social connection: %v", err)
		http.Error( w, "failed to save social connection", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
}
