help:
	just --list

# run github actions locally
act:
	act

# Parameterized build command for different projects
build app:
    @echo "Building binary for {{app}}"
    cd {{app}} && CGO_ENABLED=0 GOOS=linux go build -o bin/app

lint: go-lint ts-lint buf-lint
test: go-test ts-test

buf-lint:
	buf lint

go-checks: go-lint go-test
ts-checks: ts-lint ts-test

go-coverage-report:
	go test -coverprofile=reports/go/coverage.out ./...
	go tool cover -func reports/go/coverage.out -o reports/go/coverage.txt
	go tool cover -html reports/go/coverage.out -o reports/go/cover.html

go-lint-changed: install-go-tools
	@echo Linting recently changed go files
	golangci-lint run --fix --new-from-rev=HEAD~1 --config .golangci.yaml

go-lint: install-go-tools
	@echo Linting go files
	golangci-lint run --fix --config .golangci.yaml --timeout 5m --concurrency 4

go-integration-test: install-go-tools
	# these are only meant to be ran within tilt-ci. they require external services like mysql & minio
	go test -tags integration ./...

go-test: install-go-tools
	@echo Running go tests
	go test -short -timeout 30s ./...

ts-lint: install-ts-tools
	@echo Linting typescript
	cd ui && npm run format

ts-test: install-ts-tools
	@echo Running typescript tests
	# cd ui && npm run test

ts-build: install-ts-tools
	@echo Building typescript
	cd ui && npm run build

install-tools: install-go-tools install-ts-tools

# Run CI pipeline. This should be preferred to tilt ci as it creates and destroys the test cluster.
ci: install-tools
	ctlptl create cluster kind --registry=ctlptl-registry \
	&& \
	tilt ci \
	&& \
	ctlptl delete cluster kind

generate-openapi: install-go-tools
	@echo Generating OpenAPI
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency --parseInternal --dir api --output services/api/docs

generate-api-clients: generate-openapi generate-go-api-client generate-typescript-client generate-k6
	@echo Generated API clients

generate-k6:
	@echo Generating K6 tests
	@echo You will have to modify the generated script.js to work with your API
	docker run --rm \
		-v "./services/api/docs/swagger.json:/local/swagger.json" \
		-v "./k6/tests/openapi:/out" \
		openapitools/openapi-generator-cli generate \
		-i "/local/swagger.json" \
		-g k6 \
		-o "/out" \
		--skip-validate-spec

# rest
generate-go-api-client:
	# https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#16---docker
	@echo Generating Go API client
	docker run \
		-v "./services/api/docs/swagger.yaml:/local/swagger.yaml" \
		-v "./pkg/clients/api:/out" \
		openapitools/openapi-generator-cli \
		generate \
		-i "/local/swagger.yaml" \
		-g go \
		-o "/out/" \
		-p packageName=client \
		-p withGoMod=false

generate-typescript-client: install-ts-tools
	@echo Deprecated: generating graphql instead
	just generate-graph-ts-client

# mockery test mocks defined in .mockery.yaml
generate-go-mocks:
	mockery

# struct to struct data mapping using goverter
generate-go-converters: install-go-tools
	@echo Generating converters
	go install github.com/jmattheis/goverter/cmd/goverter@v1.5.1
	goverter gen ./internal/goverter


generate-proto:
	@echo Generating protobuf
	buf generate

go-mod-download:
	@echo Download go.mod dependencies
	go mod download

install-go-tools: go-mod-download

install-ts-tools:
	@echo Installing tools from package.json
	cd ui && npm install

setup-asdf:
	cut -d\  -f1 .tool-versions|grep -E '^[^#]'|xargs -L1 asdf plugin add
	asdf install

setup:
	pre-commit install --install-hook
	pre-commit install --hook-type commit-msg
	@echo Ensure you run "setup-asdf" or install .tool-versions manually

# regenerate the graph server using gqlgen
generate-graph-server:
	@echo Generating graph server
	go get github.com/99designs/gqlgen@v0.17.55
	go install github.com/99designs/gqlgen
	cd services/graph && gqlgen generate

# download the graphql schema used to generate clients; requires the graph server to be running
generate-graph-schema:
	@echo This requires the graph service running in tilt to generate the schema
	go get github.com/alexflint/go-arg
	go get github.com/suessflorian/gqlfetch
	cd tools/graphql-schema && go run main.go

# generate graphql clients (from schema)
generate-graph-clients: generate-graph-golang-client generate-graph-ts-client
	echo Generating graphql clients
	echo This will fail if the graph service is not running in tilt

generate-graph-golang-client: generate-graph-schema
	@echo Generating golang graphql client
	go run github.com/Khan/genqlient tools/graphql-golang-client/genqlient.yaml

generate-graph-ts-client: generate-graph-schema
	@echo Generating typescript graphql client
	@echo This requires the graph service running in tilt to generate the schema
	cd ui && npm run graphql-codegen

install-amalgam-cli:
	go install ./amalgam-cli/cmd
