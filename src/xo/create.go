package main

import "fmt"

func create(args []string) {
	fmt.Println("Creating GitHub Reposittory: " + args[2])
	github(args[2])
	fmt.Println("Creating Buildkite Pipeline: " + args[2])
	buildkite(args[2])
}
