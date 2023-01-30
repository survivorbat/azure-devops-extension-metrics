MAKEFLAGS := --no-print-directory --silent

.PHONY: example

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

protogen: ## Generate protobuf files
	protoc -I=resources resources/*.proto --ts_out=./client/src/message
	protoc -I=resources resources/*.proto --go_out=.

rs: ## Run the server locally
	cd server && LISTENER_URL="localhost:3000" PROCESSOR_ID=redis REDIS_HOST=localhost:6379 REDIS_PASSWORD="" REDIS_DB=0 go run .

bc: ## Build the client
	cd client && npm run compile

bs: ## Build the server docker image
	docker-compose -p azdometrics -f example/docker-compose.yaml build

rc: ## Run the client example locally
	cd example && node extension.js

fmt: ## Format the code
	cd server && go fmt ./... && go mod tidy
	cd client && npm run fmt

example: ## Run the example docker-compose
	docker-compose -p azdometrics -f example/docker-compose.yaml up -d
