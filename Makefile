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

# Run a worker service locally to index and structure data from a decentralized source
# into the RSS3 Protocol format. Use `make worker WORKER_ID=<worker-id>`.
# Example: `make worker WORKER_ID=ethereum-uniswap` runs one worker to index and transform data from the Uniswap dApp source.
#
# To use this command:
# - Set the decentralized network endpoint (to which the decentralized source belongs) in the 'endpoint' section of deploy/config.yaml.
# - Configure the decentralized data source component (e.g., a specific blockchain or its dApp) with the necessary parameters
#   in the 'component' section of deploy/config.yaml.
# - Use the worker ID of the component as WORKER_ID in the make command.
worker: generate
	@if [ -z "$(WORKER_ID)" ]; then \
		echo "WORKER_ID is not set. Usage: make worker WORKER_ID=<worker-id>"; \
		exit 1; \
	fi
	ENVIRONMENT=development go run ./cmd \
			--module=worker \
			--worker.id=$(WORKER_ID)


# Run the core service of the RSS3 Node locally, providing indexed data conforming to the RSS3 protocol to end users.
# The node serves API requests on port 80 by default. Use `make core`.

# Note: When routing API requests to check indexed data, include the routing group name in some paths.
# For example, API services within 'internal/node/component/decentralized' may require adding '/decentralized' to
# their paths, whereas API services within 'internal/node/component/info' may not.
core: generate
	ENVIRONMENT=development go run ./cmd \
			--module=core

# Start the monitor service to check the status of workers. Use `make monitor`.
monitor: generate
	ENVIRONMENT=development go run ./cmd \
			--module=monitor


# Start the broadcaster service to synchronize the node's status with a global indexer through heartbeats.
# Use `make broadcaster`.
#
# To use this command:
# - Generate the config.yaml for this service by registering an RSS3 node on the testnet at
#   'https://explorer.testnet.rss3.io/nodes'. This YAML file includes essential information needed to start the service.
# - Place the config.yaml file within the 'deploy' directory.
#
# Note:
# - Set the environment type to 'beta' at the top of config.yaml. The global indexer will verify the node's
#   accessibility. If the type is set to 'alpha', it will not be verified.
# - Registering an RSS3 node may require staking a certain number of RSS3 tokens (if the node is in 'Alpha' type).
broadcaster: generate
	ENVIRONMENT=development go run ./cmd \
			--module=broadcaster
