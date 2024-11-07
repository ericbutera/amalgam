.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

act: ## Run act
	act

test: install-tools ## Run tests
	@echo Running tests
	go test -v -timeout 30s ./... -short
	# TODO: typescript (ui)

# lint: install-tools ## Run linter
# 	@echo Running UI linters
# 	cd ui && npm run lint
# 	# TODO: buf lint (doesn't work in ci right now)
# 	# TODO: graphql lint

go-lint: install-tools ## Run golang linter
	@echo Running golang linters
	go vet ./...
	golangci-lint run
	staticcheck ./...

ts-lint: ## Run typescript linter
	@echo Running typescript linters
	cd ui && npm run lint

ci: install-tools ## Run CI pipeline
	ctlptl create cluster kind --registry=ctlptl-registry \
	&& \
	tilt ci \
	&& \
	ctlptl delete cluster kind

generate-openapi: install-tools  ## Generate OpenAPI spec
	@echo Generating OpenAPI
	swag init --parseDependency --parseInternal --dir api --output api/docs

generate-api-clients: generate-openapi generate-go-api-client generate-typescript-client generate-k6 ## Generate api & graphql clients
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

generate-typescript-client: ## Generate Typescript client
	@echo Generating Typescript graphql client
	cd ui && npx graphql-codegen

generate-proto: ## Generate protobuf
	@echo Generating protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	buf generate

generate-converters: ## Generate copygen type converters
	@echo Generating converters
	@go get github.com/switchupcb/copygen
	@go install github.com/switchupcb/copygen
	copygen -yml internal/copygen/setup.yml

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %
