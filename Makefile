# Makefile

GOPATH=$(PWD)/vendor:$(PWD)/src
GOBIN=$(PWD)/bin
GOFILES=$(wildcard src/*.go)
GONAME=xo
PID=/tmp/go-$(GONAME).pid
.DEFAULT_GOAL := help

.PHONY: help test build publish-release publish-docker clean

help:
	@printf "\nUSAGE: make [command] e.g. make app \n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@printf '\n'

lint: ## Lints the repository source code
	@echo "[LINT] Linting Repository Source"

test: ## Runs go test
	@echo "[TEST] Running Tests"

build: test ## Builds the Go Binaries
	@echo "[BUILD] Building application to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(GONAME) $(GOFILES)

publish-release: ## Publish to GitHub Releases
	@echo "[PUBLISH] Publishing to GitHub Releases"

publish-docker: ## Publish to DockerHub
	@echo "[PUBLISH] Publishing to DockerHub"

clean:  ## Publish to DockerHub
	@echo "[CLEAN] Cleaning Up"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean


app: lint build publish-release publish-docker  ## Lints, Builds and publishes the application
