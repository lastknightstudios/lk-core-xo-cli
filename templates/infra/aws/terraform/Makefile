# Makefile

# Variables
SHELL=/bin/bash
VARS="./deployments/$(ENTITY_NAME).$(DEPLOYMENT_ACCOUNT).tfvars"
WORKSPACE="$(ENTITY_NAME).$(DEPLOYMENT_ACCOUNT)"

.DEFAULT_GOAL := help
.PHONY: deploy help	init-env init plan apply

# Targets

help:
	@printf "\nUSAGE: make [command] e.g. make solution (ensure both ENV and REGION are exported)\n\n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@printf '\n'

init-env:
	@if [ -z $(ENITY_NAME) ]; then \
		echo "[INIT] ENITY_NAME is not set in .env file"; \
	 fi
	@if [ -z $(DEPLOYMENT_ACCOUNT) ]; then \
		echo "[INIT] ACCOUNT is not set in .env file"; \
	 fi
	@if [ -z $(REGION) ]; then \
		echo "[INIT] REGION is not set in .env file"; \
		exit 1; \
	 fi
	@if [ ! -f "$(VARS)" ]; then \
		echo "[INIT] Expected tfvars file is missing: ./$(VARS), Specify correct ENV variable. e.g. export ENV=dev.australia-southeast1.tfvars"; \
		exit 1; \
	 fi

init: init-env ## Initialise the environment and run terraform init
	@echo "[INIT] Initialising Terraform"
	@terraform init \
			-input=false \
			-force-copy \
			-lock=true \
			-upgrade \
			-verify-plugins=true
	@echo "[INIT] Selecting Terraform workspace $(WORKSPACE)"
	@terraform workspace select $(WORKSPACE) || terraform workspace new $(WORKSPACE)

lint: ## Not Yet Implemented: Lints the repository source code
	@echo "[LINT] Linting Repository Source"

test: ## Not Yet Implemented: Runs go test
	@echo "[TEST] Running Tests"

plan: init ## Run terraform plan showing expected resource deployment
	@terraform plan \
		-lock=true \
		-input=false \
		-refresh=true \
		-var-file="$(VARS)"

plan-destroy: init ## Creates a terraform destroy plan.
	@terraform plan \
		-input=false \
		-refresh=true \
		-destroy \
		-var-file="$(VARS)"

apply: init ## Run terraform apply and deploys the solution.
	@terraform apply \
		-lock=true \
		-input=false \
		-refresh=true \
		-var-file="$(VARS)" \
		-lock=false

destroy: init ## [WARNING!] Destroy all things
	@terraform destroy \
		-lock=true \
		-input=false \
		-refresh=true \
		-var-file="$(VARS)" \
		-lock=false

solution: apply ## Deploys the infrastructure
