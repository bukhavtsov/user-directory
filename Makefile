.DEFAULT_GOAL         := help
REPO                  := github.com/bukhavtsov/user-directory
BIN_PATH              ?= ./bin
SERVER_IMAGE_NAME     ?= server:latest
SERVER_CONTAINER_NAME ?= server_container
SERVER_SRC_PATH       ?= ./cmd/
SERVER_BIN_PATH       ?= $(BIN_PATH)/server/server
SERVER_DOCKER_PATH    ?= ./docker/server

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
run: build
	@echo "Running server..."
	$(SERVER_BIN_PATH)