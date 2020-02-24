package main

import (
	"fmt"
	"os"
)

func create(repository string, pipeline string, project string) {

	// Check Required ENV Vars are set
	envVarCheck("XO_REPO_ORG")
	envVarCheck("XO_REPO_TOKEN")
	envVarCheck("XO_PIPELINE_ORG")
	envVarCheck("XO_PIPELINE_TOKEN")
	os.Exit(1)

	// Load Plugins
	repo, pipe := load(repository, pipeline)
	repo.CreateRepository(project)
	pipe.CreatePipeline(project)

}

func envVarCheck(envVar string) {

	value, exists := os.LookupEnv(envVar)

	if exists {
		fmt.Println("[CONFIG]", envVar, value, "is set.")
	} else {
		fmt.Println("[ERROR] Environment Variable needs setting:", envVar)
	}
}
