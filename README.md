# lk-core-xo-cli

A CLI Tool plugin framework for provisioning new development projects and resources. Supporting multiple organisations.

## Usage

These variables must be exported to the environment.

- The XO_GITHUB_ORG is the github stub name for  your targetted org
- The XO_GITHUB_TOKEN is your personal access token.

```bash
export XO_GITHUB_ORG=lastknightstudios
export XO_GITHUB_TOKEN=abcds123123123123123
```

## Installation

* Clone this repository.
* Install prerequistes ( optional )
* Build the tool.

### Clone this repository

You will need to have git installed to follow these steps.

```bash
sudo apt install git
```

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

Building the tool will create a docker image and add xo to PATH.

```bash
make xo
```

The tool is now ready to use.

## Usage

* Configure the xo tool.
* Create a new project.

### Configure the xo tool

The xo tool requires some basic configuration before you can create a repository and pipeline.
It will require api keys for to your repository and pipeline services.

Todo: For github do this
Todo: For buildkite do this

```bash
xo config
```

Follow the prompts. Credentials are currently stored in ~/.xo/credentials.

Todo: Describe credential overrides here.

### Create a new project

Ensure you are in the folder you want the new project to clone into.
e.g. cd ~/scm

```bash
xo create project-name
```

## Branching Strategy

* `feature/*` - Feature branches. Pull requests into `master`
* `bugfix/*`  - Bugfix branches. Pull requests into `master`
* `master`    - Production Release Branch
