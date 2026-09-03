package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/google/uuid"
)

type PublisherService struct {
	store store.Querier
}

func NewPublisherService(store store.Querier) *PublisherService {
	return &PublisherService{
		store: store,
	}
}

func (s *PublisherService) CreateJobs(ctx context.Context, postID, userID uuid.UUID, socialConnectionIDs []uuid.UUID) ([]store.Publisher, error) {
	if len(socialConnectionIDs) == 0 {
		return nil, nil
	}

	publishers, err := s.store.CreatePublisherJob(ctx, store.CreatePublisherJobParams{
		PostID:              postID,
		UserID:              userID,
		SocialConnectionIds: socialConnectionIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create publisher jobs: %w", err)
	}

	return publishers, nil
}
