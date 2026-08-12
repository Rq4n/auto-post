package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type SocialHandler struct {
	socialService service.SocialService
}

func NewSocialHandler(socialService service.SocialService) *SocialHandler {
	return &SocialHandler{
		socialService: socialService,
	}
}

type AccountPayload struct {
	provider string
}

func (s *SocialHandler) handleConnectNewSocialAccount(w http.ResponseWriter, r *http.Request) {
	// Pega o user_id colocado pelo AuthMiddleware
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		log.Printf("failed to get user_id from session context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	parseID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	sID := pgtype.UUID{
		Bytes: parseID,
		Valid: true,
	}
	var payload AccountPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}
	conn, err := s.socialService.ConnectNewSocial(
		r.Context(),
		sID,
		payload.provider,
	)
	if err != nil {
		log.Printf("failed to connect account %v", err)
		http.Error(w, "failed to connect account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(conn); err != nil {
		log.Printf("failed to encode post response: %v", err)
	}
}
