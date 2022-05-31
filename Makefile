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

.PHONY: docker-down
docker-down:
	docker-compose down

.PHONY: docker-build
docker-build:
	docker-compose build

.PHONY: docker-test
docker-test:
	docker-compose run chat-service-test

.PHONY: docker-dev
docker-dev:
	docker-compose run chat-service-dev

.PHONY: docker-prod
docker-prod:
	docker-compose run chat-service-prod