package main

import (
	"fmt"
	"os"
	"plugin"
)

// Repository Interface
type Repository interface {
	CreateRepository(name string)
	CreateWebhook(webhook string)
}

// Pipeline Interface
type Pipeline interface {
	CreateRepository()
	CreateWebhook()
}

func load(repository string, pipeline string) Repository {
	return loadRepositoryPlugin(repository, "Repository")
}

func loadRepositoryPlugin(repositoryPlugin string, pluginSymbol string) Repository {
	fmt.Println("Loading:", repositoryPlugin, " plugin")

	var mod string
	mod = "./bin/" + repositoryPlugin + ".so"

	plug, err := plugin.Open(mod)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symbol, err := plug.Lookup(pluginSymbol)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var repo Repository
	repo, ok := symbol.(Repository)

	if !ok {
		fmt.Println("unexpected type from plugin symbol")
		os.Exit(1)
	}

	fmt.Println("Plugin Loaded:", repositoryPlugin, repo)

	return repo
}

func loadPipelinePlugin(plugin string) {
	fmt.Println("Loading:", plugin, " plugin")
}
