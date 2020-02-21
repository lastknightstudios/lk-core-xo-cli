package main

import "fmt"

// Plugin interfaces

type repository interface {
	url() string
	token() string
	createRepository() bool
}

type pipeline interface {
	url() string
	token() string
	createPipeline() bool
}

func create(args []string) {

	fmt.Println("Creating Reposittory: " + args[2])
	//github(args[2])
	fmt.Println("Creating Pipeline: " + args[2])
	//buildkite(args[2])
}
