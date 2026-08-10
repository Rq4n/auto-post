package service

import (
	"context"

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

func (s *SocialService) ConnectNewSocial(ctx context.Context, provider string) (*store.Post, error) {
	return nil, nil
}
