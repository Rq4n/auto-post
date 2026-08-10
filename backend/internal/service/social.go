package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
)

type SocialService struct {
	store store.Querier
}

func NewSocialService(store store.Querier) *SocialService {
	return &SocialService{
		store: store,
	}
}

func (s *SocialService) ConnectNewSocial(ctx context.Context, provider string) (*store.SocialConnection, error) {
	social, err := s.store.ConnectNewProvider(ctx, provider)
	if err != nil {
		return nil, fmt.Errorf("failed to start provider %w", err)
	}
	return &social, nil
}
