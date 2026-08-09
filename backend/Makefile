include .env
export

goose-up:
	@goose -dir db/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE)" up

goose-down:
	@goose -dir db/migrations postgres "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSLMODE)" down

goose-create:
	@goose -dir db/migrations create -s $(name) sql

build:
	@go build -o bin/autopost ./cmd/api/

run: build
	@./bin/autopost
