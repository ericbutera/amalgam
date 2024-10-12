.DEFAULT_GOAL := help

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: generate-api-client
generate-api-client: ## Generate API client
	# TODO: find a cleaner way to do this
	# TODO: regenerate api docs (cd api && make docs) before running this command
	#
	# https://github.com/OpenAPITools/openapi-generator?tab=readme-ov-file#16---docker
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