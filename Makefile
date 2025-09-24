
.PHONY: help build build-dev build-prod up-dev up-dev-detached down-dev restart-dev exec-dev up-prod down-prod restart-prod exec-prod logs logs-tail logs-dev logs-dev-tail clean clean-all gen-test-cover-svg backend-test

help:
	@echo "Available commands:"
	@echo "  build            - Build production Docker image"
	@echo "  build-dev        - Build development Docker image"
	@echo "  build-prod       - Build production Docker image (alias for build)"
	@echo ""
	@echo "  up-dev           - Start development environment"
	@echo "  up-dev-detached  - Start development environment in background"
	@echo "  down-dev         - Stop development environment"
	@echo "  restart-dev      - Restart development environment"
	@echo "  exec-dev         - Execute bash in development container"
	@echo ""
	@echo "  up-prod          - Start production environment"
	@echo "  down-prod        - Stop production environment"
	@echo "  restart-prod     - Restart production environment"
	@echo "  exec-prod        - Execute bash in production container"
	@echo ""
	@echo "  logs             - View production logs"
	@echo "  logs-tail        - View production logs (follow)"
	@echo "  logs-dev         - View development logs"
	@echo "  logs-dev-tail    - View development logs (follow)"
	@echo ""
	@echo "  clean            - Clean Docker system and volumes"
	@echo "  clean-all        - Clean all Docker data"
	@echo ""
	@echo "  backend-test     - Run Go tests"
	@echo "  gen-test-cover-svg - Generate test coverage SVG"

build:
	docker-compose --env-file .env.prod build --no-cache

build-dev:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev build --no-cache

build-prod:
	docker-compose --env-file .env.prod build --no-cache

up-dev:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev up --build

up-dev-detached:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev up --build -d

down-dev:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev down

restart-dev:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev restart

exec-dev:
	docker exec -it ghostbin_backend_dev bash

up-prod:
	docker-compose --env-file .env.prod up -d --build --force-recreate

down-prod:
	docker-compose --env-file .env.prod down

restart-prod:
	docker-compose --env-file .env.prod restart

exec-prod:
	docker exec -it ghostbin_backend bash

logs:
	docker-compose --env-file .env.prod logs

logs-tail:
	docker-compose --env-file .env.prod logs -f

logs-dev:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev logs

logs-dev-tail:
	docker-compose -f docker-compose.dev.yml --env-file .env.dev logs -f

clean:
	docker system prune -f
	docker volume prune -f

clean-all:
	docker system prune -af
	docker volume prune -f

gen-test-cover-svg:
	go install github.com/nikolaydubina/go-cover-treemap@latest
	go test -coverprofile cover.out ./...
	go-cover-treemap -coverprofile cover.out > ./assets/testcover.svg
	rm ./cover.out

backend-test:
	go test -v ./...
