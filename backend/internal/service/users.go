package service

import (
	"context"
	"fmt"

	"github.com/Rq4n/autopost/internal/store"
)

type UserService struct {
	store store.Querier
}

func NewUserService(store store.Querier) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) CreateUser(ctx context.Context, Email, GoogleID string) (*store.User, error) {
	user, err := s.store.CreateUser(ctx, store.CreateUserParams{
		GoogleID: GoogleID,
		Email:    Email,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user %w", err)
	}
	return &user, nil
}

func (s *UserService) GetUserByGoogleID(ctx context.Context, googleID string) (*store.User, error) {
	user, err := s.store.GetUserByGoogleID(ctx, googleID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// buscar social_connections do usuario no banco
// join tables  social_connections -> user.id -> provider

	
