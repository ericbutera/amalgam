help:
	just --list

# run github actions locally
act:
	act

# Parameterized build command for different projects
build app:
    @echo "Building binary for {{app}}"
    cd {{app}} && CGO_ENABLED=0 GOOS=linux go build -o bin/app

install-amalgam-cli:
	go install ./amalgam-cli/cmd

lint: go-lint ts-lint buf-lint
test: go-test ts-test

buf-lint:
	buf lint

go-checks: go-lint go-test
ts-checks: ts-lint ts-test

go-generate-mocks:
	go install github.com/vektra/mockery/v2@v2.46.3
	mockery

go-coverage-report:
	go test -coverprofile=reports/coverage.out ./...
	go tool cover -func reports/coverage.out -o reports/coverage.txt
	go tool cover -html reports/coverage.out -o reports/cover.html

go-lint-changed: install-go-tools
	@echo Linting recently changed go files
	golangci-lint run --fix --new-from-rev=HEAD~1 --config .golangci.yaml

go-lint: install-go-tools
	@echo Linting go files
	golangci-lint run --fix --config .golangci.yaml --timeout 5m --concurrency 4

go-integration-test: install-go-tools
	# these are only meant to be ran within tilt-ci. they require external services like mysql & minio
	go test -v -tags integration ./...

go-test: install-go-tools
	@echo Running go tests
	go test -short -timeout 30s ./...

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

# Generate api & graphql clients
generate-api-clients: generate-openapi generate-go-api-client generate-typescript-client generate-k6
	@echo Generated API clients

# Generate tests from OpenAPI spec
generate-k6:
	@echo Generating K6 tests
	@echo You will have to modify the generated script.js to work with your API
	docker run --rm \
		-v "./api/docs/swagger.json:/local/swagger.json" \
		-v "./k6/tests/openapi:/out" \
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
		-v "./pkg/clients/api:/out" \
		openapitools/openapi-generator-cli \
		generate \
		-i "/local/swagger.yaml" \
		-g go \
		-o "/out/" \
		-p packageName=client \
		-p withGoMod=false

# Generate Typescript client
generate-typescript-client: install-ts-tools
	@echo Deprecated: generating graphql instead
	just generate-graph-ts-client


# Generate protocol buffers with buf
generate-proto:
	@echo Generating protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	buf generate

# Generate copygen type converters
generate-converters: install-go-tools
	@echo Generating converters
	go get github.com/switchupcb/copygen
	go install github.com/switchupcb/copygen
	copygen -yml internal/copygen/setup.yml

go-mod-download:
	@echo Download go.mod dependencies
	go mod download

install-go-tools: go-mod-download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0

install-ts-tools:
	@echo Installing tools from package.json
	cd ui && npm install

setup:
	pre-commit install --install-hook
	pre-commit install --hook-type commit-msg

# Generate graph
generate-graph-server:
	@echo Generating graph server
	go get github.com/99designs/gqlgen@v0.17.55
	go install github.com/99designs/gqlgen
	cd graph && gqlgen generate

# Generate the gql client schema
generate-graph-schema:
	# TODO: run service, generate schema, produce artifact
	@echo This requires the graph service running in tilt to generate the schema
	go get github.com/alexflint/go-arg
	go get github.com/suessflorian/gqlfetch
	cd tools/graphql-schema && go run main.go

# Generate the golang graph client
generate-graph-golang-client: generate-graph-schema
	# TODO: use artifact created from `generate-graph-schema` not ://service/query
	@echo Generating golang graphql client
	go run github.com/Khan/genqlient tools/graphql-golang-client/genqlient.yaml


generate-graph-ts-client: generate-graph-schema
	# TODO: use artifact created from `generate-graph-schema` not ://service/query
	@echo Generating typescript graphql client
	@echo This requires the graph service running in tilt to generate the schema
	cd ui && npm run graphql-codegen
