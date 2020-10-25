# xo cli tool - alpha * use at own risk

**DISCLAIMER**: Some of this readme maybe inaccurate or incomplete. This statement will be removed when merged into master.

## Contents

- Description & Motivation
- Building the xo tool
- Using the xo tool
- Contributing & Creating Plugins
- References

---

A CLI Tool plugin framework for provisioning new development projects and resources. Supporting multiple organisations and accounts. Uses SSH Tokens for api access the tool enforces coding patterns and standards by deploying an optional opinionated cookie cutter template based on selected project type on initial commit. These cookie cutters can be customised.

The tool creates the repositories, pipeline and webhooks between the tools.

## Motivation

My main motivations in developing this tool is to learn the go programming language and is to prevent clickops of common tasks by providing a simple cli tool that enforces standard patterns by means of good generic cookie cutter templates.

## Building the xo tool

You can either build binaries or create a docker container. Building the tool will create the go binaries and store in the ./bin folder or create a small docker image < 25MB.

### Secrets

These secrets must be exported to the environment for the xo tool to be able to read them in. How you manage your secrets is out of scope of this README. I personally read them in from AWS Secrets Manager by calling a custom shell profile function when i need them. In a pipeline these are imported via AWS CLI and exported to environment then cleared after use.

- The XO_REPO_ORG is the github stub name for  your targeted org or username
- The XO_REPO_TOKEN is an ssh key that has access to create and commit to repos. *
- The XO_PIPELINE_ORG is the ci/cd stub name for your targeted org or username
- The XO_PIPELINE_TOKEN is your api key or access token that has access to create and execute pipelines. *

```bash
export XO_REPO_ORG=lastknightstudios
export XO_REPO_TOKEN=abcds123123123123123
export XO_PIPE_ORG=lastknightstudios
export XO_PIPE_TOKEN=abcds123123123123123
```

#### Tool Requirements

This solution uses the following tools

- git
- docker
- make

#### Clone this repository

Clone the xo repository

For HTTPS

```bash
git clone https://github.com/lastknightstudios/lk-core-xo-cli.git
```

For SSH

```bash
git clone git@github.com:lastknightstudios/lk-core-xo-cli.git
```

Change into the xo directory

```bash  
cd lk-core-xo-cli
```

#### Make

The make is self documenting. Just run make the default goal is help.

```bash
$ make

USAGE: make [command] e.g. make app

all                  Lint, Test and Build and Publish
app                  Builds the xo go app
build-all            Build both the Go App and the Docker Image
clean                Runs go clean
docker               Builds the Docker Image
lint                 Lints the repository source code
publish-all          Publishes the application to container repo and github releases
publish-dockerrepo   Publish to dockerrepo
publish-release      Publish to GitHub Releases
test                 Runs go test
```

#### Binary Builder

```bash
make app
```

#### Docker Builder

```bash
make docker
```

---

### Using the xo tool

As the tool is provided by means of binaries or docker image the method of use is described below

#### Binary Usage

If using the binaries ensure you have updated your path to point to the project bin folder.

```bash
export PATH=$(pwd)/bin:$PATH
```

or install the binary into your environment using your prefered method.

Change to the folder you want the new project to clone into

```bash
cd ~/scm/github
```

To create a repository and pipeline

```bash
xo create my_project --repo github --pipeline buildkite
```

To create a repository and pipeline with commit webhook

```bash
xo create -h my_project --repo github --pipeline buildkite
```

To create just a repository

```bash
xo create --project my_project --repo github
```

To create just a pipeline

```bash
xo create my_project --pipeline buildkite
```

#### Docker Usage

```bash
docker run lastknight/xo create my_project --repo github --pipeline buildkite
```

---

### Contributing & Creating Plugins

The xo tool utilises interfaces that must be implemented for them to be correctly loaded. Some variables and return types must be utilised. The implementation is up to you.

#### Creating Repository Plugins

Create a new plugin file in src/plugins directory copy the following code in to the plugins directory

```go
package main

import (
    // ...
)

type repository string

var repoOrg = os.Getenv("XO_REPO_ORG")
var repoToken = os.Getenv("XO_REPO_TOKEN")

func (g repository) CreateRepository(name string) {
 _CreateRepository(name)
}

func (g repository) CreateWebhook(webhook string) {
 _CreateWebhook(webhook)
}

// Repository exported as symbol
var Repository repository

// Plugin Implementation

func _CreateRepository(name string) {
 // Your implementation

}

func _CreateWebhook(webhook string) {
 // Your implementation
}

```

or create a skeleton template automatically with the tool

``` bash
xo plugin bitbucket --type repository
```

#### Creating Pipeline Plugins

Create a new plugin file in src/plugins directory copy the following code in to the plugins directory

```go
package main

import (
// ...
)

type pipeline string

var pipelineOrg = os.Getenv("XO_PIPELINE_ORG")
var pipelineToken = os.Getenv("XO_PIPELINE_TOKEN")

func (g pipeline) CreatePipeline(name string) {
 _CreatePipeline(name)
}

// Pipeline exported as symbol
var Pipeline pipeline

// Plugin Implementation

func _CreatePipeline(name string) {
    // Your Implementation
}
```

or create a skeleton template automatically with the tool

``` bash
xo plugin travisci --type pipeline
```

##### Branching Strategy

The master branch is protected. Feel free to create feature or bug branchs and issue a pull request.

- `feature/*` - Feature branches. Pull requests into `master`
- `bugfix/*`  - Bugfix branches. Pull requests into `master`
- `master`    - Production Release Branch
