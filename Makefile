.DEFAULT_GOAL := help

PROJECT_NAME := GIGMSN Subscriber

.PHONY: help
help:
	@echo "------------------------------------------------------------------------"
	@echo "${PROJECT_NAME}"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z0-9_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run tests
	@docker-compose up --build test
	@docker-compose rm -fsv test

.PHONY: consumer/up
consumer/up: ## run consumer container
	@docker-compose up --build consumer

.PHONY: consumer/stop
consumer/stop: ## stop and remove consumer container
	@docker-compose rm -fsv consumer
