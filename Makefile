.PHONY: vendor
vendor:
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test-local
test-local:
	go test -mod=vendor ./...

.PHONY: cover
cover:
	go test -coverprofile="coverage.out" ./... && go tool cover -html="coverage.out"

.PHONY: down
down:
	docker-compose down

.PHONY: build
build:
	docker-compose build

.PHONY: test
test:
	docker-compose up chat-service-test

.PHONY: dev
dev:
	docker-compose up chat-service-dev

.PHONY: prod
prod:
	docker-compose up chat-service-prod