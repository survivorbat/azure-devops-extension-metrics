MAKEFLAGS := --no-print-directory --silent

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

protogen: ## Generate protobuf files
	protoc -I=resources resources/*.proto --ts_out=./client/src/message
	protoc -I=resources resources/*.proto --go_out=.

rs: ## Run the server
	cd server && LISTENER_URL="localhost:8080" go run .

rc: ## Run the client example
	cd client && npm run compile
	cd example && node extension.js

fmt: ## Format the code
	cd client && npm run fmt
	cd server && go fmt ./...
