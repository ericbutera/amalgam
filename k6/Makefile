.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate
generate: ## Generate tests from OpenAPI spec
	# TODO: regenerate OpenAPI spec before running (cd ../api && make docs)
	cp ../api/docs/swagger.json . \
	&& \
	docker run --rm \
		-v "${PWD}:/local" \
		-v "./tests:/out" \
		openapitools/openapi-generator-cli generate \
		-i "/local/swagger.json" \
		-g k6 \
		-o "/out" \
		--skip-validate-spec \
	&& \
	rm swagger.json

.PHONY: test
test:  ## Run tests
	# k6 run tests/test.js
	docker run --rm \
		-v "./tests:/tests" \
		grafana/k6 \
		run \
		"/tests/script.js"
