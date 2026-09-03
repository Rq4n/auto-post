package handler

import (
	"github.com/Rq4n/autopost/internal/service"
	"github.com/Rq4n/autopost/internal/store"
)

type Handler struct {
	handlePost      *PostHandler
	handleUser      *UserHandler
	handleTwitter   *TwitterHandler
	handlePublisher *PublisherHandler
}

func NewHandler(repo store.Querier) *Handler {
	userService := service.NewUserService(repo)
	publisherService := service.NewPublisherService(repo)
	postService := service.NewPostService(repo, publisherService)
	socialService := service.NewSocialService(repo)

	return &Handler{
		handlePost:      NewPostsHandler(postService),
		handleUser:      NewUserHandler(*userService),
		handlePublisher: NewPublisherHandler(*publisherService),
		handleTwitter:   NewTwitterHandler(*socialService),
	}
}
