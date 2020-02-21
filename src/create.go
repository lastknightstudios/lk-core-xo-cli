package main

import "fmt"

// Plugin interfaces

type repository interface {
	name() string
	url() string
	token() string
}

type pipeline interface {
	name() string
	url() string
	token() string
}

func create(args []string) {

	fmt.Println("Creating Reposittory: " + args[2])
	//github(args[2])
	fmt.Println("Creating Pipeline: " + args[2])
	//buildkite(args[2])
}

func createRepository(repository repository) {
	fmt.Println(repository)
}

func createPipeline(pipeline pipeline) {
	fmt.Println(pipeline)
}
