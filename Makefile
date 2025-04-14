# Makefile

# Load environment variables from .env file
include .env
export

# Database URL construction
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

# Migrations
migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-create:
	migrate create -ext sql -dir migrations -seq $(DB_NAME)

migrate-status:
	@echo "Migrations status:"
	migrate -path migrations -database "$(DB_URL)" version

db-reset:
	dropdb $(DB_NAME) || true
	createdb $(DB_NAME)

db-tables:
	@echo "\nListe des tables dans la base de données:"
	@psql -d $(DB_NAME) -c "\dt"

run:
	@echo "Chargement des variables d'environnement..."
	@source .env && go run cmd/server/main.go

.PHONY: migrate-up migrate-down migrate-create migrate-status db-reset db-tables