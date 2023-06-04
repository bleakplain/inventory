.PHONY: all build clean test

APP_NAME := inventory-service
GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean
GOTEST := $(GO) test
WIRE := wire
DOCKER := docker

all: build

build: wire
	$(GOBUILD) -o $(APP_NAME) ./cmd/main.go

clean:
	$(GOCLEAN)
	rm -f $(APP_NAME)

test:
	$(GOTEST) -v ./...

wire:
	$(WIRE) ./internal/infrastructure/wire.go

docker-build:
	$(DOCKER) build -t $(APP_NAME) .

docker-run: docker-build
	$(DOCKER) run -p 8080:8080 -p 9000:9000 $(APP_NAME)

