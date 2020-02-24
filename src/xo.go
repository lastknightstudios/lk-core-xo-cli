package main

import (
	"fmt"
	"os"
)

// Secrets Global
var Secrets map[string]string

// Config Global
var Config map[string]string

func main() {

	// Init Globals
	Secrets = make(map[string]string)
	Config = make(map[string]string)

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {

	case "version":
		fmt.Println(version())
		os.Exit(0)

	case "help":
		help()
		os.Exit(0)

	case "create":
		if len(os.Args) < 4 {
			fmt.Println("[ERROR] Missing all required arguments")
			help()
			os.Exit(1)
		}

		Config["project"] = os.Args[2]
		Config["repository"] = os.Args[3]
		Config["pipeline"] = os.Args[4]

		// Gather Secrets
		getEnvVars()

		// Create Resources
		create(Config["repository"], Config["pipeline"], Config["project"])

	default:
		fmt.Println("invalid command:", "try xo help")
	}
}
