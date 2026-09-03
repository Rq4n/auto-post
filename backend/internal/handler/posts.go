package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
	"github.com/google/uuid"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostsHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

type PostsPayload struct {
	Title                string      `json:"title"`
	Body                 string      `json:"content"`
	SocialConnectionsIDs []uuid.UUID `json:"social_connections_ids"`
}

func (p *PostHandler) CreateAndPublish(w http.ResponseWriter, r *http.Request) {
	// 1. Extração do UserID do contexto
	userID, ok := r.Context().Value("userID").(uuid.UUID)
	if !ok {
		log.Printf("failed to get user_id from session context")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// 2. Decode do Payload
	var payload PostsPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload format or invalid UUIDs", http.StatusBadRequest)
		return
	}

	// 3. Validação dos campos obrigatórios
	if payload.Title == "" || payload.Body == "" {
		http.Error(w, "title and content are required", http.StatusBadRequest)
		return
	}

	// 4. Chamada do serviço passando todos os parâmetros necessários
	post, publishers, err := p.postService.CreateAndPublishPost(
		r.Context(),
		userID,
		payload.SocialConnectionsIDs,
		payload.Title,
		payload.Body,
	)
	if err != nil {
		log.Printf("failed to create post: %v", err)
		http.Error(w, "failed to create post", http.StatusInternalServerError)
		return
	}

	// 5. Montagem da resposta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := struct {
		Post       any `json:"post"`
		Publishers any `json:"publishers"`
	}{
		Post:       post,
		Publishers: publishers,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("failed to encode post response: %v", err)
	}
}
