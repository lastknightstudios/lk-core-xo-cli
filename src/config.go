package main

import (
	"fmt"
	"os"
)

func getEnvVars() {
	Config["repoOrg"] = envVarCheck("XO_REPO_ORG")
	Secrets["repoToken"] = envVarCheck("XO_REPO_TOKEN")
	Config["pipelineOrg"] = envVarCheck("XO_PIPELINE_ORG")
	Secrets["pipelineToken"] = envVarCheck("XO_PIPELINE_TOKEN")
	// TODO: Tokens should be validated before use
}

func envVarCheck(envVar string) string {
	value, exists := os.LookupEnv(envVar)

	if exists {
		fmt.Println("[CONFIG] Environment variable:", envVar+" loaded.")
	} else {
		fmt.Println("[ERROR] Environment variable needs setting:", envVar)
		// Not happy with this hard exit i think all values should be checked first and shown which are not set.
		os.Exit(1)
	}

	return value
}
