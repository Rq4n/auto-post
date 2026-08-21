package handler

import (
	"net/http"

	"github.com/Rq4n/autopost/internal/service"
)

type PublisherHandler struct {
	publisherService service.PublisherService
}

func NewPublisherHandler(publisherService service.PublisherService) *PublisherHandler {
	return &PublisherHandler{
		publisherService: publisherService,
	}
}

func (s *PublisherHandler) CreatePublisher(w http.ResponseWriter, r *http.Request) {
	// if err := json.NewDecoder(r.Body).Decode(r); err != nil {
	// 	http.Error(w, "invalid payload", http.StatusBadRequest)
	// 	return
	// }
}
