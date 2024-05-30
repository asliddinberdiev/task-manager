.SILENT:
-include .env

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Default targets
.PHONY: help run print build up down migrate-up migrate-down

help:
	@echo "Makefile commands:"
	@echo "make run              - Golang application and bot run"
	@echo "make print            - Print db url"
	@echo "make build     		   - Build the Docker image for the application"
	@echo "make up        		   - Start all services of the application"
	@echo "make down      		   - Stop all services of the application"
	@echo "make migrate-up       - Run database migrations up"
	@echo "make migrate-down     - Run database migrations down"

run:
	@go run app/cmd/main.go
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
	@migrate -path app/migrations -database "$(DB_URL)" -verbose up

migrate-down:
	@migrate -path app/migrations -database "$(DB_URL)" -verbose down


.PHONY: all
