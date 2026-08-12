package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/jackc/pgx/v5/pgtype"
)

type SocialService struct {
	store store.Querier
}

func NewSocialService(store store.Querier) *SocialService {
	return &SocialService{
		store: store,
	}
}

func (s *SocialService) ConnectNewSocial(ctx context.Context, userID pgtype.UUID, provider string) (*store.SocialConnection, error) {
	arg := store.ConnectNewProviderParams{
		UserID:   userID,
		Provider: provider,
	}
	social, err := s.store.ConnectNewProvider(ctx, arg)
	if err != nil {
		log.Printf("failed to connect new provider %v", err)
		return nil, fmt.Errorf("failed to start provider %w", err)
	}
	return &social, nil
}
