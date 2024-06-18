VERSION=$(shell git describe --tags --abbrev=0)
COMMIT=$(shell git rev-parse --short HEAD)

# Build this project and generate the binary in the ./build directory
.PHONY: build
build: generate
	mkdir -p ./build
	go build \
		-ldflags "-X github.com/rss3-network/node/internal/constant.Version=$(VERSION) -X github.com/rss3-network/node/internal/constant.Commit=$(COMMIT)" \
		-o ./build/node ./cmd
DOCKER_COMPOSE_FILE=./.devcontainer/docker-compose.yaml

ifeq ($(VERSION),)
	VERSION="0.0.0"
endif

# Run gocyclo to check the cyclomatic complexities of the top 10 most complex functions
gocyclo:
	go run github.com/fzipp/gocyclo/cmd/gocyclo@latest -top 10 -ignore "_test|contract_" ./

# Run gocognit to check cognitive complexities of the top 10 most complex functions
gocognit:
	go run github.com/uudashr/gocognit/cmd/gocognit@latest -top 10 -ignore "_test|contract_" ./

# Sort struct fields to reduce memory usage
align:
	go run github.com/dkorunic/betteralign/cmd/betteralign@latest -apply ./...

# Run go generate to process any //go:generate directives
generate:
	go generate ./...

# Format the codebase by running gocognit and align
.PHONY: fmt
fmt: gocognit align

# Run golangci-lint to lint the whole codebase, ensuring code quality
lint:
	go mod tidy
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1 run

# Run all tests with coverage and enabled race detection
test:
	go test -cover -race -v ./...

# Start all docker services defined in the docker-compose file (located in the ./devcontainer directory)
.PHONY: service_up
service_up:
	docker compose --project-name rss3-node -f $(DOCKER_COMPOSE_FILE) up -d

# Stop and remove all docker services defined in the docker-compose file (located in the ./devcontainer directory)
.PHONY: service_down
service_down:
	docker compose --project-name rss3-node -f $(DOCKER_COMPOSE_FILE) down

# Function to start a docker service defined in the docker-compose file
define start_service
	docker compose --project-name rss3-node -f $(DOCKER_COMPOSE_FILE) up -d $(1)
endef

# Function to stop a docker service defined in the docker-compose file
define stop_service
	docker compose --project-name rss3-node -f $(DOCKER_COMPOSE_FILE) stop $(1)
endef

# Function to remove a docker service defined in the docker-compose file
define remove_service
	docker compose --project-name rss3-node -f $(DOCKER_COMPOSE_FILE) rm -f $(1)
endef

# Allow users to start a specific service by passing the SERVICE variable
# Example: `make docker-start SERVICE=redis` Start the specific redis service in docker
.PHONY: docker-start
docker-start:
	@if [ -z "$(SERVICE)" ]; then \
		echo "SERVICE is not set. Usage: make docker-start SERVICE=<service-name>"; \
		exit 1; \
	fi
	$(call start_service,$(SERVICE))

# Allow users to stop a specific service by passing the SERVICE variable
# Example: `make docker-stop SERVICE=redis` Stop the specific redis service in docker
.PHONY: docker-stop
docker-stop:
	@if [ -z "$(SERVICE)" ]; then \
		echo "SERVICE is not set. Usage: make docker-stop SERVICE=<service-name>"; \
		exit 1; \
	fi
	$(call stop_service,$(SERVICE))

# Allow users to remove a specific service by passing the SERVICE variable
# Example: `make docker-remove SERVICE=redis` Remove the specific redis service in docker
.PHONY: docker-remove
docker-remove:
	@if [ -z "$(SERVICE)" ]; then \
		echo "SERVICE is not set. Usage: make docker-remove SERVICE=<service-name>"; \
		exit 1; \
	fi
	$(call stop_service,$(SERVICE))
	$(call remove_service,$(SERVICE))

# Build a Docker image for this project
image: generate
	docker build \
    		--tag rss3-network/node:$(VERSION) \
    		.




# The following section details the make commands used to run various node services,
# including worker, core, monitor, and broadcaster service.
#
# To run any services, you must have certain Docker services running, such as redis and cockroachdb
# Use `make service_up` to start all required services using docker-compose
# Use `make service_down` to remove all required services
#
# Run a worker service locally to index and structure data from a decentralized source
# into the RSS3 Protocol format. Use `make worker WORKER_ID=<worker-id>`
worker: generate
	@if [ -z "$(WORKER_ID)" ]; then \
		echo "WORKER_ID is not set. Usage: make worker WORKER_ID=<worker-id>"; \
		exit 1; \
	fi
	ENVIRONMENT=development go run ./cmd \
			--module=worker \
			--worker.id=$(WORKER_ID)

# Run the core service of the RSS3 Node locally, providing indexed data conforming to the RSS3 protocol to end users
core: generate
	ENVIRONMENT=development go run ./cmd \
			--module=core

# Start the monitor service to check the status of workers. Use `make monitor`
monitor: generate
	ENVIRONMENT=development go run ./cmd \
			--module=monitor

# Start the broadcaster service to synchronize the node's status with a global indexer through heartbeats
# Use `make broadcaster`
broadcaster: generate
	ENVIRONMENT=development go run ./cmd \
			--module=broadcaster
