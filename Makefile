# Makefile

APPNAME=xo
DOCKERREPO=lastknightstudios
GOPATH=$(PWD)/vendor:$(PWD)/src
GOBIN=$(PWD)/bin
GOFILES=$(wildcard src/*.go)
GONAME=$(APPNAME)
PID=/tmp/go-$(GONAME).pid
.DEFAULT_GOAL := help
.PHONY: help test build publish-release publish-dockerrepo clean

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
	@echo "[BUILD] Building Docker Image"
	@docker build --no-cache --pull . -t $(DOCKERREPO)/$(APPNAME):latest

publish-release: ## Publish to GitHub Releases
	@echo "[PUBLISH] Publishing to GitHub Releases"

publish-dockerrepo: ## Publish to dockerrepo
	@echo "[PUBLISH] Publishing to dockerrepo"
	@docker login docker.io -u $(DOCKERREPO)
	@docker push $(DOCKERREPO)/$(APPNAME):latest
	@docker logout

clean:  ## Publish to dockerrepo
	@echo "[CLEAN] Cleaning Up"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

app: lint build clean ## Lint, Test and Build

publish-all: publish-release publish-dockerrepo ## Publishes the application to container repo and github releases