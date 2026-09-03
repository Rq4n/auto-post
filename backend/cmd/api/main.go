// Package main
package main

import (
	"fmt"
	"log"

	"github.com/Rq4n/autopost/internal/auth"
	"github.com/Rq4n/autopost/internal/config"
	"github.com/Rq4n/autopost/internal/handler"
	"github.com/Rq4n/autopost/internal/service"
	"github.com/Rq4n/autopost/internal/store"
	"github.com/Rq4n/autopost/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load config file ")
	}

	auth.NewAuth(
		cfg.Providers.Google,
		cfg.Providers.Twitter,
	)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DB,
		cfg.DB.SSLMode,
	)

	dbConfig := db.Config{
		DSN:          dsn,
		MaxIdleTime:  cfg.DB.MaxIdleTime,
		MaxIdleConns: cfg.DB.MaxOpenConn,
		MinIdleConns: cfg.DB.MinIdleConn,
	}

	pool, err := db.StartConnection(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	repo := store.New(pool)

	defer pool.Close()
	log.Print("connection pool established")

	app := &app{
		addr: ":8080",
		dbconfig: dbconfig{
			dbx: pool,
			dbc: &dbConfig,
		},
		Handler: Handler{
			handleUser:      *handler.NewUserHandler(*service.NewUserService(repo)),
			handlePublisher: *handler.NewPublisherHandler(*service.NewPublisherService(repo)),
			handleTwitter:   *handler.NewTwitterHandler(*service.NewSocialService(repo)),
			handlePost:      *handler.NewPostsHandler(service.NewPostService(repo, service.NewPublisherService(repo))),
		},
	}

	mux := app.mount()
	log.Fatal(app.start(mux))
}
