export SHELL:=/bin/bash
export BASE_NAME:=$(shell basename ${PWD})

help: ## Prints help for targets with comments
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}'
	@echo ""

format-install: ## Install go formatting dependencies on the localhost
	@go install golang.org/x/tools/cmd/goimports@latest \
	 	&& go install mvdan.cc/gofumpt@latest

format: ## Format the code
	@go mod tidy \
		&& goimports -w . \
		&& gofumpt -l -w .

test: ## Run Test
	@go test ./...