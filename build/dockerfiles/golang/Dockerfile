# ------------------------------------------------------------------------------
# Base image
# ------------------------------------------------------------------------------
FROM golang:1.18-bullseye AS golang
RUN apt-get update && apt-get install git
ARG UID=1000
ARG GID=1000
ENV UID=${UID}
ENV GID=${GID}
RUN getent passwd $UID || (groupadd --gid $GID user && useradd --uid $UID --gid user --shell /bin/bash --create-home user)
USER $UID
WORKDIR $GOPATH/src
COPY . $GOPATH/src
ENV GO111MODULE="on" \
    GOOS=linux

# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang AS golang_test
WORKDIR $GOPATH/src
ENTRYPOINT go test -v ./...

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang AS golang_dev
ENTRYPOINT go run cmd/main.go

# ------------------------------------------------------------------------------
# Build image
# ------------------------------------------------------------------------------
FROM golang AS golang_build
RUN go build -mod=vendor -o $GOPATH/bin/main $GOPATH/src/cmd/main.go
RUN chmod +x $GOPATH/bin/main


# ------------------------------------------------------------------------------
# Production image
# ------------------------------------------------------------------------------
FROM alpine:3.7 as golang_prod
COPY --from=golang_build /go/bin/ ./
EXPOSE ${APP_PORT}
CMD ["./main"]