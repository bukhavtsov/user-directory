.DEFAULT_GOAL         := help
REPO                  := github.com/bukhavtsov/user-directory
BIN_PATH              ?= ./bin
SERVER_IMAGE_NAME     ?= server:latest
SERVER_CONTAINER_NAME ?= server
SERVER_SRC_PATH       ?= ./cmd/
SERVER_BIN_PATH       ?= $(BIN_PATH)/server/server
SERVER_DOCKER_PATH    ?= ./docker/server

DB_IMAGE_NAME     ?= db-users:latest
DB_CONTAINER_NAME ?= db-users
DB_SRC_PATH       ?= ./docker/db-users
DB_DOCKER_PATH    ?= $(DB_SRC_PATH)/Dockerfile

PHONY: help
help: ## makefile targets description
	@echo "Usage:"
	@egrep '^[a-zA-Z_-]+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##/#-/' | column -t -s "#"

.PHONY: fmt
fmt: ## automatically formats Go source code
	@echo "Running 'go fmt ...'"
	@go fmt -x "$(REPO)/..."

.PHONY: build
build: fmt ## compile package and dependencies
	@echo "Building server..."
	CGO_ENABLED=0 go build -o $(SERVER_BIN_PATH) $(SERVER_SRC_PATH)

.PHONY: run
run: build ## execute server binary
	@echo "Running server..."
	$(SERVER_BIN_PATH)

.PHONY: image
image: build ## build image from Dockerfile ./docker/server/Dockerfile
	@echo "Building server image..."
	cp $(SERVER_BIN_PATH) $(SERVER_DOCKER_PATH)
	@docker build -t $(SERVER_IMAGE_NAME) $(SERVER_DOCKER_PATH)
	rm $(SERVER_DOCKER_PATH)/server
	@echo "Building db image..."
	@docker build -f $(DB_DOCKER_PATH) -t $(DB_IMAGE_NAME) .

.PHONY: up
up : image ## up docker compose
	@docker-compose up