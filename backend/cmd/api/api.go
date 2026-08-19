// Package
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Rq4n/autopost/internal/auth"
	"github.com/Rq4n/autopost/pkg/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type app struct {
	addr string
	dbconfig
	Handler
}
type dbconfig struct {
	dbx *pgxpool.Pool
	dbc *db.Config
}

type Handler struct {
	handlePost      *PostHandler
	handleUser      *UserHandler
	handleSocial    *SocialHandler
	handlePublisher *PublisherHandler
}

func (a *app) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte("status: ok"))
	})

	r.Route("/v1", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Post("/post", a.handlePost.handleCreateNewPost) // create the post (title, content), may vary depending on social media

		// should handle the provider selection such as twitter etc.
		// frontend checkbox selects providers with checkbox and make GET request /v1/provider 
		r.Post("/social", a.handleSocial.handleConnectNewProvider) 
	})

	r.Get("/auth/{provider}", beginAuthProviderCallback)
	r.Get("/auth/{provider}/callback", a.handleUser.getAuthCallbackFunction)

	return r
}

func (a *app) start(mux http.Handler) error {
	srv := &http.Server{
		Addr:    a.addr,
		Handler: mux,
	}

	log.Printf("server started on port: %v", a.addr)

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		log.Print("signal caught", "signal", s.String())
		shutdown <- srv.Shutdown(ctx)
	}()

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	log.Print("server has stopped")
	return nil
}
