package main

import (
	"fmt"
	"os"
)

func main() {
	// Parse Arguments

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "version":
		fmt.Println(version())
		os.Exit(1)
	case "help":
		help()
		os.Exit(1)
	case "create":
		if len(os.Args) < 4 {
			fmt.Println("[ERROR] Missing all required arguments")
			help()
			os.Exit(1)
		}
		project := os.Args[2]
		repository := os.Args[3]
		pipeline := os.Args[4]
		create(repository, pipeline, project)
	default:
		fmt.Println("invalid command:", "try xo help")
	}

}
