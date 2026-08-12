package main

import (
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
)

type SocialHandler struct {
	socialService service.SocialService
}

func NewSocialHandler(socialService service.SocialService) *SocialHandler {
	return &SocialHandler{
		socialService: socialService,
	}
}

func (s *SocialHandler) ConnectNewSocialAccount(w http.ResponseWriter, r *http.Request) {
}
