.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## Generate graph
	@echo "Generating graph server"
	gqlgen generate
	@echo "Generating supporting go code"
	go generate ./...


.PHONY: build
build: ## Build the binary
	CGO_ENABLED=0 GOOS=linux go build -o bin/app
