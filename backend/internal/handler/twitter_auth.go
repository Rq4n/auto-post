package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/auth"
	"github.com/Rq4n/autopost/internal/service"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

func (t *TwitterHandler) GetProviderByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	val := ctx.Value(auth.UserIDKey)
	if val == nil {
		log.Printf("failed to get user from context")
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := val.(uuid.UUID)
	if !ok {
		log.Printf("invalid uuid format in context")
		http.Error(w, "invalid id", http.StatusInternalServerError)
		return
	}

	providers, err := t.socialService.GetProviderByUserID(ctx, userID)
	if err != nil {
		http.Error(w, "Erro ao buscar redes sociais conectadas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(providers); err != nil {
		http.Error(w, "Erro ao codificar resposta JSON", http.StatusInternalServerError)
		return
	}
}

func (t *TwitterHandler) BeginProviderAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// 1. Extrai o userID do contexto da requisição
	userID, ok := r.Context().Value(auth.UserIDKey).(uuid.UUID)
	if !ok {
		log.Printf("failed to get user_id from context in BeginProviderAuth")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var gothProvider string
	switch provider {
	case "twitter":
		gothProvider = "twitterv2"
	default:
		http.Error(w, "unsupported provider", http.StatusBadRequest)
		return
	}

	// 2. Anexa o userID no parâmetro 'state' da URL para o OAuth preservar no redirecionamento
	q := r.URL.Query()
	q.Set("state", userID.String())
	r.URL.RawQuery = q.Encode()

	// O Goth exige a string literal "provider" no contexto para funcionar
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

	// Passa a string "provider" para o Goth reconhecer o driver
	ctx := context.WithValue(r.Context(), "provider", gothProvider)
	r = r.WithContext(ctx)

	pvUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("failed to complete %s oauth: %v", gothProvider, err)
		http.Error(w, "failed to complete oauth", http.StatusInternalServerError)
		return
	}

	// Extrai o UserID tipado como uuid.UUID usando a chave centralizada auth.UserIDKey
	userID, ok := r.Context().Value(auth.UserIDKey).(uuid.UUID)
	if !ok {
		log.Printf("failed to get user_id from context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	_, err = t.socialService.ConnectNewProvider(r.Context(), userID, provider, pvUser)
	if err != nil {
		log.Printf("failed to create social connection: %v", err)
		http.Error(w, "failed to save social connection", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
}
