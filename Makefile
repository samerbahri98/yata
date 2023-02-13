include .env
export

default: dev

.PHONY: build
build:
	go build

.PHONY: sqlc
sqlc:
	docker run --rm -v $(shell pwd)/internal/model:/src -w /src kjconroy/sqlc generate ;

.PHONY: generate-migration
goose-generate:
	sh ./scripts/generate-migration $n

.PHONY: generate-entity
generate-entity:
	sh ./scripts/generate-entity.sh $n

.PHONY: goose-migrate
goose-migrate:
	sh ./scripts/goose-migrate.sh $f

.PHONY: dev
dev:
	docker compose -f ./configs/docker/docker-compose.yml up ;

.PHONY: dev-down
dev-down:
	docker compose -f ./configs/docker/docker-compose.yml down -v ;

.PHONY: lint
lint:
	docker run --rm -v $(shell pwd):/src -w /src golangci/golangci-lint:v1.51.1 golangci-lint run ;

.PHONY: sqlite
sqlite:
	sh ./scripts/sqlite.sh
