.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

act: ## Run act
	act --pull=false -P ubuntu-latest=catthehacker/ubuntu:act-latest -W .github/workflows/all.yaml

test: install-tools ## Run tests
	@echo Running tests
	go test -v ./... -short

lint: install-tools ## Run linter
	@echo Running linter
	go vet ./... && \
	golangci-lint run && \
	staticcheck ./...

ci: install-tools test lint ## Run CI pipeline
	ctlptl create cluster kind --registry=ctlptl-registry \
	&& \
	tilt ci \
	&& \
	ctlptl delete cluster kind

generate-openapi: install-tools  ## Generate OpenAPI spec
	@echo Generating OpenAPI
	swag init --parseDependency --parseInternal --dir api --output api/docs

generate-api-clients: generate-openapi generate-go-api-client generate-typescript-api-client generate-k6 ## Generate API clients from OpenAPI spec
	@echo "Generated API clients"

generate-k6: ## Generate tests from OpenAPI spec
	@echo Generating K6 tests
	docker run --rm \
		-v "./api/docs/swagger.json:/local/swagger.json" \
		-v "./k6/tests:/out" \
		openapitools/openapi-generator-cli generate \
		-i "/local/swagger.json" \
		-g k6 \
		-o "/out" \
		--skip-validate-spec

generate-go-api-client: ## Generate Golang API client
	# TODO: regenerate OpenAPI spec before running (cd api && make docs)
	# https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#16---docker
	@echo Generating Go API client
	docker run \
		-v "./api/docs/swagger.yaml:/local/swagger.yaml" \
		-v "./pkg/client:/out" \
		openapitools/openapi-generator-cli \
		generate \
		-i "/local/swagger.yaml" \
		-g go \
		-o "/out/" \
		-p packageName=client \
		-p withGoMod=false

generate-typescript-api-client: ## Generate Typescript API client
	# TODO: regenerate OpenAPI spec before running (cd api && make docs)
	# https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#16---docker
	@echo Generating Typescript API client
	docker run \
		-v "./api/docs/swagger.yaml:/local/swagger.yaml" \
		-v "./ui/app/lib/client:/out" \
		openapitools/openapi-generator-cli \
		generate \
		-i "/local/swagger.yaml" \
		-g typescript-fetch \
		-o "/out/"

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %