.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

generate-server: ## Generate graph
	@echo Generating graph server
	gqlgen generate

generate-gqlclient-schema: ## Generate the gql client schema
	go get github.com/suessflorian/gqlfetch
	go run generate/client/main.go

generate-go-client: ## Generate the golang graph client
	@echo Generating graph client
	go run github.com/Khan/genqlient generate/client/genqlient.yaml
	go generate ./...

.PHONY: build
build: ## Build the binary
	CGO_ENABLED=0 GOOS=linux go build -o bin/app
