# Makefile

APPNAME=xo
SHELL := /bin/bash
DOCKERREPO ?= lastknightstudios
REPO ?= github
PIPELINE ?= buildkite
CGO_ENABLED=0
GOOS=linux
GOPATH=$(PWD)/vendor:$(PWD)/src
GOBIN=$(PWD)/bin
GOFILES=$(wildcard src/*.go)
GOPLUGINS="$(wildcard src/plugins/*.go)
GONAME=$(APPNAME)
PID=/tmp/go-$(GONAME).pid

.DEFAULT_GOAL := help
.PHONY: help test app docker publish-release publish-dockerrepo build-all publish-all clean

help:
	@printf "\nUSAGE: make [command] e.g. make app \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@printf '\n'

lint: ## Not Yet Implemented: Lints the repository source code
	@echo "[LINT] Linting Repository Source"

test: ## Not Yet Implemented: Runs go test
	@echo "[TEST] Running Tests"

build-all: app docker ## Build both the Go App and the Docker Image

app: test ## Builds the Go App !!! not ideal
	@echo "[BUILD] Building plugins to ./bin"
	@go build -buildmode=plugin -o bin/buildkite.so src/plugins/buildkite.go
	@go build -buildmode=plugin -o bin/github.so src/plugins/github.go
	@echo "[BUILD] Building application to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build -a -installsuffix nocgo -o bin/$(GONAME) $(GOFILES)

docker: ## Builds the Docker Image
	@echo "[BUILD] Building Docker Image"
	@docker build --no-cache --pull . -t $(DOCKERREPO)/$(APPNAME):latest
	
publish-all: publish-release publish-dockerrepo ## Publishes the application to container repo and github releases

publish-release: ## Not Yet Implemented: Publish to GitHub Releases
	@echo "[PUBLISH] Publishing to GitHub Releases"

publish-dockerrepo: ## Publish to dockerrepo
	@echo "[PUBLISH] Publishing to dockerrepo"
	@docker login docker.io -u $(DOCKERREPO)
	@docker push $(DOCKERREPO)/$(APPNAME):latest
	@docker logout

clean:  ## Runs go clean
	@echo "[CLEAN] Cleaning Up"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

all: lint test build-all publish-all clean run ## Lint, Test and Build and Publish
