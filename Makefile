.SILENT:
-include .env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Default targets
.PHONY: help run print build up down migrate-up migrate-down

help:
	@echo "Makefile commands:"
	@echo "make app-run          - Golang application run"
	@echo "make bot-run          - Bot run"
	@echo "make print            - Print db url"
	@echo "make build     		   - Build the Docker image for the application"
	@echo "make up        		   - Start all services of the application"
	@echo "make down      		   - Stop all services of the application"
	@echo "make migrate-up       - Run database migrations up"
	@echo "make migrate-down     - Run database migrations down"

app-run:
	@go run app/cmd/main.go

bot-run:
	@nodemon bot/bot.js

print:
	@echo $(DB_URL)

build:
	@docker-compose build

up:
	@docker-compose up -d

down:
	@docker-compose down

migrate-up:
	@migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	@migrate -path migrations -database "$(DB_URL)" -verbose down


.PHONY: all
