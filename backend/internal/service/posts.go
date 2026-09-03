package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/google/uuid"
)

type PostService struct {
	store            store.Querier
	publisherService *PublisherService // new service injection
}

func NewPostService(store store.Querier, publisherService *PublisherService) *PostService {
	return &PostService{
		store:            store,
		publisherService: publisherService,
	}
}

func (s *PostService) CreateAndPublishPost(
	ctx context.Context,
	userID uuid.UUID,
	socialConnectionIDs []uuid.UUID,
	title, body string,
) (*store.Post, []store.Publisher, error) {
	post, err := s.store.CreateNewPosts(ctx, store.CreateNewPostsParams{UserID: userID, Title: title, Body: body})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create post: %w", err)
	}

	publishers, err := s.publisherService.CreateJobs(ctx, post.ID, userID, socialConnectionIDs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to publish post: %w", err)
	}

	return &post, publishers, nil
}
