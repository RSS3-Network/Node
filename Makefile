VERSION=$(shell git describe --tags --abbrev=0)
COMMIT=$(shell git rev-parse --short HEAD)

ifeq ($(VERSION),)
	VERSION="0.0.0"
endif

generate:
	go generate ./...

lint: generate
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run

test:
	go test -cover -race -v ./...

.PHONY: build
build: generate
	mkdir -p ./build
	go build \
		-ldflags "-X github.com/rss3-network/node/internal/constant.Version=$(VERSION) -X github.com/rss3-network/node/internal/constant.Commit=$(COMMIT)" \
		-o ./build/node ./cmd

image: generate
	docker build \
    		--tag rss3-network/node:$(VERSION) \
    		.

run: generate
	ENVIRONMENT=development go run ./cmd
