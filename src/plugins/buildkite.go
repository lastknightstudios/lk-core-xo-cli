package main

import "fmt"

type pipeline string

func (g pipeline) CreatePipeline() {
	fmt.Println("Creating Buildkite Pipeline")
}

// Pipeline exported as symbol
var Pipeline pipeline
