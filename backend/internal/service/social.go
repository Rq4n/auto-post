package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Rq4n/autopost/internal/store"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/markbates/goth"
)

type SocialService struct {
	store store.Querier
}

func NewSocialService(store store.Querier) *SocialService {
	return &SocialService{
		store: store,
	}
}

func (s *SocialService) ConnectNewProvider(
	ctx context.Context,
	userID uuid.UUID,
	provider string,
	pvUser goth.User,
) (*store.SocialConnection, error) {
	arg := store.ConnectNewProviderParams{
		UserID:         userID,
		Provider:       provider,
		ProviderUserID: pvUser.UserID,
		AccessToken:    pvUser.AccessToken,
		RefreshToken: pgtype.Text{
			String: pvUser.RefreshToken,
			Valid:  pvUser.RefreshToken != "",
		},
		ExpiresAt: pgtype.Timestamptz{
			Time:  pvUser.ExpiresAt,
			Valid: pvUser.ExpiresAt.IsZero(),
		},
	}

	social, err := s.store.ConnectNewProvider(ctx, arg)
	if err != nil {
		log.Printf("failed to create social connection: %v", err)
		return nil, fmt.Errorf("failed to create social connection: %w", err)
	}

	return &social, nil
}

func (s *SocialService) GetProviderByUserID(ctx context.Context, userID uuid.UUID) ([]store.GetProviderByUserIDRow, error) {
	provider, err := s.store.GetProviderByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get socials %w", err)
	}

	return provider, nil
}
