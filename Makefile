VERSION=$(shell git describe --tags --abbrev=0)
COMMIT=$(shell git rev-parse --short HEAD)

ifeq ($(VERSION),)
	VERSION="0.0.0"
endif

gocyclo:
	go run github.com/fzipp/gocyclo/cmd/gocyclo@latest -top 10 -ignore "_test|contract_" ./

gocognit:
	go run github.com/uudashr/gocognit/cmd/gocognit@latest -top 10 -ignore "_test|contract_" ./

# sort struct fields to reduce memory usage
align:
	go run github.com/dkorunic/betteralign/cmd/betteralign@latest -apply ./...

generate:
	go generate ./...

lint:
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1 run

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
