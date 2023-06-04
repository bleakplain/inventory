.PHONY: all build clean test wire docker

all: build

build: wire
	go build -o ./bin/inventory-service ./cmd/main.go

clean:
	rm -rf ./bin

test:
	go test -v ./...

wire:
	wire ./internal/wire.go

docker:
	docker build -t inventory-service .

run:
	./bin/inventory-service -conf ./config.yaml

