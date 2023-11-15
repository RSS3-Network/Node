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
	go test -cover -v ./...

.PHONY: build
build: generate
	mkdir -p ./build
	go build \
		-ldflags "-X github.com/naturalselectionlabs/rss3-node/internal/constant.Version=$(VERSION) -X github.com/naturalselectionlabs/rss3-node/internal/constant.Commit=$(COMMIT)" \
		-o ./build/rss3-node ./cmd

image: generate
	docker build \
    		--tag naturalselectionlabs/rss3-node:$(VERSION) \
    		.

run: generate
	ENVIRONMENT=development go run ./cmd
