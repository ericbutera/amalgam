.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## Generate tests from OpenAPI spec
	@echo Deprecated: use "cd .. && make generate-k6" instead

.PHONY: test
test:  ## Run tests
	# k6 run tests/test.js
	docker run --rm \
		-v "./tests:/tests" \
		grafana/k6 \
		run \
		"/tests/script.js"
