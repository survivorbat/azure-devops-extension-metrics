MAKEFLAGS := --no-print-directory --silent

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

protogen: ## Generate protobuf files
	protoc -I=resources resources/*.proto --ts_out=./package/src/message
	protoc -I=resources resources/*.proto --go_out=.

rs: ## Run the server
	cd listener && LISTENER_URL="localhost:8080" go run .

rc: ## Run the client example
	cd package && npm run compile
	cd example && node extension.js

fmt: ## Format the code
	cd package && npm run fmt
	cd listener && go fmt ./...
