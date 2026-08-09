package main

import (
	"encoding/json"
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostsHandler(postsService service.PostService) *PostHandler {
	return &PostHandler{
		postService: postsService,
	}
}

type PostsPayload struct {
	Title   string
	Content string
}

func (p *PostHandler) handleCreateNewPost(w http.ResponseWriter, r *http.Request) {
	var payload PostsPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	post, err := p.postService.CreateNewPost(r.Context(), payload.Title, payload.Content)
	if err != nil {
		http.Error(w, "failed to create post", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(post)
}
