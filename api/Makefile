.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the binary
	# TODO: remove GOARCH
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app

.PHONY: docs
docs:
	# TODO: remove
	@echo Deprecated: use "cd .. && make generate-openapi" instead
	@exit 1

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: image-build
image-build: ## Build docker image
	# TODO: remove
	docker build -t ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG} .

.PHONY: image-push
 image-push: ## Push docker image
	# TODO: remove
	docker push ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: run
run: ## Run app
	go run ${MAIN} server

.PHONY: test
test: ## Run tests
	go test -timeout 30s ./...