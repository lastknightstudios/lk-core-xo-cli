# Makefile. Self Documenting just run make

CGO_ENABLED=0 
GOOS=linux
APPNAME=hello-go
SHELL=/bin/bash
DOCKERREPO=lastknightstudios
GOPATH=$(PWD)/vendor:$(PWD)/app
GOBIN=$(PWD)/bin
GOFILES=$(wildcard app/*.go)
GONAME=$(APPNAME)
PID=/tmp/go-$(GONAME).pid
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
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -a -installsuffix nocgo -o bin/$(GONAME) $(GOFILES)

docker: ## Builds the Docker Image
	@echo "[BUILD] Building Docker Image"
	@docker build --no-cache --pull  . -t $(DOCKERREPO)/$(APPNAME):latest
	
publish-docker: ## Publish to dockerrepo
	@echo "[PUBLISH] Publishing to dockerrepo"
	@docker login docker.io -u $(DOCKERREPO)
	@docker push $(DOCKERREPO)/$(APPNAME):latest
	@docker logout

clean:  ## Runs go clean
	@echo "[CLEAN] Cleaning Up"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

all: lint test app docker publish-docker clean ## Lint, Test and Build and Publish