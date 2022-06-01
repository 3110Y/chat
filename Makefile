.PHONY: vendor
vendor:
	go mod vendor

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
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
	docker-compose run chat-service-test

.PHONY: dev
dev:
	docker-compose run chat-service-dev

.PHONY: prod
prod:
	docker-compose run chat-service-prod