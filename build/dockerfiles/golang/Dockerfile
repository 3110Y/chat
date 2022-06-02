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
# Build image no optimizations
# ------------------------------------------------------------------------------
FROM golang AS golang_build_no_optimizations
RUN go build -mod=vendor -o $GOPATH/bin/main_no_optimizations $GOPATH/src/cmd/main.go
RUN chmod +x $GOPATH/bin/main_no_optimizations

# ------------------------------------------------------------------------------
# Test image
# ------------------------------------------------------------------------------
FROM golang_build_no_optimizations AS golang_test
WORKDIR $GOPATH/src
ENTRYPOINT go test -v ./...

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------
FROM golang_build_no_optimizations AS golang_dev
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3
EXPOSE 3000
ENTRYPOINT $GOPATH/bin/dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec $GOPATH/bin/main_no_optimizations

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
EXPOSE 3000
ENTRYPOINT ["./main"]