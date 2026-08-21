package handler

import (
	"github.com/Rq4n/autopost/internal/service"
	"github.com/Rq4n/autopost/internal/store"
)

type Handler struct {
	handlePost      *PostHandler
	handleUser      *UserHandler
	handleSocial    *SocialHandler
	handlePublisher *PublisherHandler
}

func NewHandler(repo store.Querier) *Handler {
	postService := service.NewPostService(repo)
	userService := service.NewUserService(repo)
	socialService := service.NewSocialService(repo)
	publisherService := service.NewPublisherService(repo)

	return &Handler{
		handlePost:      NewPostsHandler(*postService),
		handleUser:      NewUserHandler(*userService),
		handleSocial:    NewSocialHandler(*socialService),
		handlePublisher: NewPublisherHandler(*publisherService),
	}
}
