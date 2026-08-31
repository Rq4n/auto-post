package service

import (
	"context"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PublisherService struct {
	store store.Querier
}

func NewPublisherService(store store.Querier) *PublisherService {
	return &PublisherService{
		store: store,
	}
}

func (s *PublisherService) CreatePublisherJob(ctx context.Context, postID, userID uuid.UUID) ([]store.Publisher, error) {
	pb, err := s.store.CreatePublisherJob(ctx, store.CreatePublisherJobParams{
		PostID: pgtype.UUID{Bytes: postID, Valid: true},
		UserID: pgtype.UUID{Bytes: userID, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return pb, nil
}
