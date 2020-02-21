package main

import (
	"fmt"
	"os"
	"plugin"
)

// Repository Interface
type Repository interface {
	Repo()
}

func load(args []string) {
	// determine plugin to load. this should be a config element
	repo := "bitbucket"

	var mod string
	switch repo {
	case "bitbucket1":

		mod = "./bin/bitbucket.so"
	case "github":
		mod = "./bin/github.so"
	default:
		fmt.Println("Repo not currently supported.")
		os.Exit(1)
	}

	// load module
	// 1. open the so file to load the symbols
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symGreeter, err := plug.Lookup("Repository")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var repository Repository
	repository, ok := symGreeter.(Repository)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	repository.Repo()

}
