help:
	just --list

# run github actions locally
act:
	act

lint: go-lint ts-lint buf-lint
test: go-test ts-test

buf-lint:
	buf lint

go-checks: go-lint go-test
ts-checks: ts-lint ts-test

go-lint: install-go-tools
	@echo Linting go
	pre-commit run golangci-lint || true
	pre-commit run go-vet || true
	pre-commit run go-staticcheck-mod || true

go-test: install-go-tools
	@echo Running go tests
	go test -timeout 30s ./...

ts-lint: install-ts-tools
	@echo Linting typescript
	cd ui && npm run lint

ts-test: install-ts-tools
	@echo Running typescript tests
	# cd ui && npm run test

ts-build: install-ts-tools
	@echo Building typescript
	cd ui && npm run build

install-tools: install-go-tools install-ts-tools

# Run CI pipeline
ci: install-tools
	ctlptl create cluster kind --registry=ctlptl-registry \
	&& \
	tilt ci \
	&& \
	ctlptl delete cluster kind

# Generate OpenAPI spec
generate-openapi: install-go-tools
	@echo Generating OpenAPI
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency --parseInternal --dir api --output api/docs

generate-api-clients: generate-openapi generate-go-api-client generate-typescript-client generate-k6 ## Generate api & graphql clients
	@echo Generated API clients

# Generate tests from OpenAPI spec
generate-k6:
	@echo Generating K6 tests
	docker run --rm \
		-v "./api/docs/swagger.json:/local/swagger.json" \
		-v "./k6/tests:/out" \
		openapitools/openapi-generator-cli generate \
		-i "/local/swagger.json" \
		-g k6 \
		-o "/out" \
		--skip-validate-spec

# Generate Golang API client
generate-go-api-client:
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

# Generate Typescript client
generate-typescript-client: install-ts-tools
	@echo Generating Typescript graphql client
	cd ui && npx graphql-codegen

# Generate protocol buffers with buf
generate-proto:
	@echo Generating protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	buf generate

# Generate copygen type converters
generate-converters: install-go-tools
	@echo Generating converters
	@go get github.com/switchupcb/copygen
	@go install github.com/switchupcb/copygen
	copygen -yml internal/copygen/setup.yml

go-mod-download:
	@echo Download go.mod dependencies
	go mod download

install-go-tools: go-mod-download
	@echo Installing tools from tools.go
	# cat tools/tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %

install-ts-tools:
	@echo Installing tools from package.json
	cd ui && npm install

setup:
	pre-commit install --install-hook
