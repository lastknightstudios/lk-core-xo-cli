# Makefile. Self Documenting just run make

APPNAME=hello-c
SRC=$(PWD)/app
CBIN=$(PWD)/bin
CFILES=$(wildcard $(SRC)/*.c)
CFLAGS=-I. -static
SHELL=/bin/bash
DOCKERREPO=lastknightstudios

.DEFAULT_GOAL := help
.PHONY: help test lint app docker publish-docker clean

help:
	@printf "\nUSAGE: make [command] e.g. make app \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@printf '\n'

lint: ## Not Yet Implemented: Lints the repository source code
	@echo "[LINT] Linting Repository Source"

test: ## Not Yet Implemented: Runs go test
	@echo "[TEST] Running Tests"

app: lint test ## Builds the Go App !!! not ideal
	@echo "[BUILD] Building application to ./bin"
	@gcc -o $(CBIN)/$(APPNAME) $(CFILES) $(CFLAGS)

docker: ## Builds the Docker Image
	@echo "[BUILD] Building Docker Image"
	@docker build --no-cache --pull  . -t $(DOCKERREPO)/$(APPNAME):latest
	
publish-docker: ## Publish to dockerrepo
	@echo "[PUBLISH] Publishing to dockerrepo"
	@docker login docker.io -u $(DOCKERREPO)
	@docker push $(DOCKERREPO)/$(APPNAME):latest
	@docker logout

all: lint test app docker publish-docker ## Lint, Test and Build and Publish