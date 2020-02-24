package main

import (
	"fmt"
	"os"
)

func create(repository string, pipeline string, project string) {

	// Check Required ENV Vars are set
	fmt.Println("[CONFIG] Checking environment variables")
	// FI: Improve this look up golang arrays and how to use them and also to return all unset vars rather than exit on the first.
	envVarCheck("XO_REPO_ORG")
	envVarCheck("XO_REPO_TOKEN")
	envVarCheck("XO_PIPELINE_ORG")
	envVarCheck("XO_PIPELINE_TOKEN")

	// TODO: Validate tokens

	// Load Plugins
	repo, pipe := load(repository, pipeline)

	// FI: Add --fix command if one of the resources has been deleted.
	repo.CreateRepository(project)
	pipe.CreatePipeline(project)

}

func envVarCheck(envVar string) {

	value, exists := os.LookupEnv(envVar)

	if exists {
		fmt.Println("[CONFIG] Environment variable:", envVar+"="+value)
	} else {
		fmt.Println("[ERROR] Environment variable needs setting:", envVar)
		os.Exit(1)
	}
}
