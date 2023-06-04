.PHONY: all build clean test wire docker

APP_NAME := inventory-service
GO := go
WIRE := wire
DOCKER := docker

all: build

build: wire
	$(GO) build -o $(APP_NAME) ./cmd

clean:
	$(GO) clean
	rm -f $(APP_NAME)

test:
	$(GO) test -v -race ./...

wire:
	$(WIRE) ./internal/wire

docker:
	$(DOCKER) build -t $(APP_NAME) .

