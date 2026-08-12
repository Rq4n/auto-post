// Package service
package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostService struct {
	store store.Querier
}

func NewPostService(store store.Querier) *PostService {
	return &PostService{
		store: store,
	}
}

func (s *PostService) CreateNewPost(ctx context.Context, userID pgtype.UUID, title string, content string) (*store.Post, error) {
	arg := store.CreateNewPostsParams{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	post, err := s.store.CreateNewPosts(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("failed to create post %w", err)
	}
	return &post, nil
}
