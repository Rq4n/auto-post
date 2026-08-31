package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostsHandler(postService service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

type PostsPayload struct {
	Title                string   `json:"title"`
	Content              string   `json:"content"`
	SocialConnectionsIDs []string `json:"social_connections_ids"`
}

func (p *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
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

	var payload PostsPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Validação básica
	if payload.Title == "" || payload.Content == "" {
		http.Error(w, "title and content are required", http.StatusBadRequest)
		return
	}

	// Cria o post associado ao usuário autenticado
	post, err := p.postService.CreateNewPost(
		r.Context(),
		sID,
		payload.Title,
		payload.Content,
	)
	if err != nil {
		log.Printf("failed to create post: %v", err)
		http.Error(w, "failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(post); err != nil {
		log.Printf("failed to encode post response: %v", err)
	}
}
