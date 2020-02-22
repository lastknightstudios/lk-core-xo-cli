package main

import (
	"fmt"
	"os"
)

func main() {

	// Parse Arguments
	//args := os.Args

	if len(os.Args) < 3 {
		help()
		os.Exit(1)
	}

	command := os.Args[1]
	project := os.Args[2]
	repository := os.Args[3]
	pipeline := os.Args[4]

	switch command {
	case "version":
		version()
	case "help":
		help()
	case "create":
		create(repository, pipeline, project)
	default:
		fmt.Println("invalid command:", "try xo help")
	}
}
