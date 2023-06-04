.PHONY: all build clean test wire docker

all: build

build:
	go build -o ./bin/inventory-service ./cmd/main.go

clean:
	rm -rf ./bin

test:
	go test -v ./...

wire:
	wire ./internal/wire/wire.go

docker:
	docker build -t inventory-service .

config:
	cp config.yaml.example config.yaml

run: build
	./bin/inventory-service -conf=config.yaml

