
build:
	docker-compose build --no-cache

up-dev:
	docker-compose -f docker-compose.dev.yml up

down-dev:
	docker-compose -f docker-compose.dev.yml down

restart-dev:
	docker-compose -f docker-compose.dev.yml restart

exec-dev:
	docker exec -it ghostbin_backend_dev bash

up-prod:
	docker-compose up -d --build --force-recreate

down-prod:
	docker-compose down

restart-prod:
	docker-compose restart

exec-prod:
	docker exec -it ghostbin_backend bash

logs:
	docker-compose logs

logs-tail:
	docker-compose logs -f

gen-test-cover-svg:
	go install github.com/nikolaydubina/go-cover-treemap@latest
	go test -coverprofile cover.out ./...
	go-cover-treemap -coverprofile cover.out > ./assets/testcover.svg
	rm ./cover.out

backend-test:
	go test -v ./...
