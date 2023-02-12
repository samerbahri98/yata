include .env
export

default: build

.PHONY: build
build:
	go build

.PHONY: sqlc
sqlc:
	docker run --rm -v $(pwd):/src -w /src kjconroy/sqlc generate

.PHONY: dev
dev:
	docker compose -f ./configs/docker/docker-compose.yml up


