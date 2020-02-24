package main

import (
	"fmt"
)

func create(repository string, pipeline string, project string) {

	// Load Plugins
	repo, pipe := load(repository, pipeline)

	// FI: Add --fix command if one of the resources has been deleted.
	repo.CreateRepository(project)
	fmt.Println("[REPO] Repository created:", project)
	pipe.CreatePipeline(project)
	fmt.Println("[PIPELINE] Pipeline created:", project)
}
