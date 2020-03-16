# Makefile

include .env

# Variables

SHELL=/bin/bash
ENVFILE=.env

GOPATH=$(PWD)/vendor:$(PWD)/src
GOBIN=$(PWD)/bin
GOFILES=$(wildcard src/*.go)
GONAME=$(APPNAME)

.DEFAULT_GOAL := help
.PHONY: help lint test app

# Targets

help:
	@printf "\nUSAGE: make [command] e.g. make app \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@printf '\n'

init-env: ## Imports .env
	@echo "[INT] Importing .env file"
	@export $($(SHELL) sed 's/=.*//' $(ENVFILE))

lint: ## Not Yet Implemented: Lints the repository source code
	@echo "[LINT] Linting Repository Source"

test: ## Not Yet Implemented: Runs go test
	@echo "[TEST] Running Tests"

app: init-env lint test ## Builds the Go App to ./bin
	@echo "[BUILD] Building application to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(GONAME) $(GOFILES)
