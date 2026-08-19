package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
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

func (s *PublisherService) FetchPendingJobs(ctx context.Context) (*store.Publisher, error) {
	_, err := s.store.FetchPendingJobs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start new publisher %w", err)
	}
	return nil, nil
}

func (s *PublisherService) UpdateJobAsProcessing(ctx context.Context, id pgtype.UUID) (*store.Publisher, error) {
	// _, err := s.store.UpdateJobAsProcessing(ctx, id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to update job as processing %w", err)
	// }

	return nil, nil
}
