# lk-core-xo-cli

A CLI Tool plugin framework for provisioning new development projects and resources. Supporting multiple organisations and accounts. Uses Personal Acccess Tokens for api access the tool enforces coding patterns and standards.
The tool creates repositories, pipelines and webhooks and also allows an initial template commit with pipeline and starter code.

## Secrets

These secrets must be exported to the environment for the xo tool to be able to read them in. How you manage your secrets is out of scope of this document. I personally read them in from AWS Secrets Manager by calling a custom shell profile function when i need them.

- The XO_REPO_ORG is the github stub name for  your targeted org or username
- The XO_REPO_TOKEN is your personal access token.
- The XO_PIPELINE_ORG is the ci/cd stub name for your targeted org or username
- The XO_PIPELINE_TOKEN is your api key or access token.

```bash
export XO_REPO_ORG=lastknightstudios
export XO_REPO_TOKEN=abcds123123123123123
export XO_PIPELINE_ORG=lastknightstudios
export XO_PIPELINE_TOKEN=abcds123123123123123
```

## Installation

- Tool Requirements
- Clone this repository.
- Build the tool.

## Tool Requirements

This solution uses the following tools

- git
- docker
- make

### Clone this repository

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

### Build the tool

You can either build binaries or create a docker container. Building the tool will create the go binaries and store in the ./bin folder or create a tiny docker image.

#### Make

The make is self documenting. Just run make the target goal is help.

```bash
$ make

USAGE: make [command] e.g. make app

all                  Lint, Test and Build and Publish
app                  Builds the Go App
build-all            Build both the Go App and the Docker Image
clean                Runs go clean
docker               Builds the Docker Image
lint                 Lints the repository source code
publish-all          Publishes the application to container repo and github releases
publish-dockerrepo   Publish to dockerrepo
publish-release      Publish to GitHub Releases
test                 Runs go test
xo                   Runs just go build
```

#### Binary Builder

```bash
make app
```

#### Docker Builder

```bash
make docker
```

### Usage

- TODO: Configure the xo tool. Will add in project config file
- Create a new project.

### Create

Ensure you are in the folder you want the new project to clone into.
e.g. cd ~/scm/project_name

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

## Branching Strategy

The master branch is protected. Feel free to create feature or bug branchs and issue a pull request.

- `feature/*` - Feature branches. Pull requests into `master`
- `bugfix/*`  - Bugfix branches. Pull requests into `master`
- `master`    - Production Release Branch
