package main

import (
	"fmt"
	"os"
	"plugin"
	"strings"
)

// Repository Interface
type Repository interface {
	CreateRepository(name string)
	CreateWebhook(webhook string)
}

// Pipeline Interface
type Pipeline interface {
	CreatePipeline(name string)
}

func load(repository string, pipeline string) (Repository, Pipeline) {
	return loadRepositoryPlugin(repository, "Repository"), loadPipelinePlugin(pipeline, "Pipeline")

}

func loadRepositoryPlugin(repositoryPlugin string, pluginSymbol string) Repository {

	var repo Repository
	sysPaths := strings.Split(envVarCheck("PATH"), ":")

	for index, sysPath := range sysPaths {

		var mod string
		mod = sysPath + "/" + repositoryPlugin + ".so"

		plug, err := plugin.Open(mod)

		if err != nil {
			fmt.Println("[PLUGINS] Discovering Repository install path, testing: ", index, sysPath)
		} else {

			symbol, err := plug.Lookup(pluginSymbol)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			repo, ok := symbol.(Repository)

			if !ok {
				fmt.Println("[ERROR] Unexpected repository type from plugin symbol")
				os.Exit(1)
			} else {
				fmt.Println("[PLUGINS] Install discovered: ", sysPath)
				fmt.Println("[PLUGINS] Plugin Loaded:", repositoryPlugin, repo)
				return repo
				break
			}
		}

	}
	return repo
}

func loadPipelinePlugin(pipelinePlugin string, pluginSymbol string) Pipeline {

	var pipe Pipeline
	sysPaths := strings.Split(envVarCheck("PATH"), ":")

	for index, sysPath := range sysPaths {

		var mod string
		mod = sysPath + "/" + pipelinePlugin + ".so"

		plug, err := plugin.Open(mod)

		if err != nil {
			fmt.Println("[PLUGINS] Discovering Pipeline install path, testing: ", index, sysPath)
		} else {

			symbol, err := plug.Lookup(pluginSymbol)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			pipe, ok := symbol.(Pipeline)

			if !ok {
				fmt.Println("[ERROR] Unexpected plugin type from plugin symbol")
				os.Exit(1)
			} else {
				fmt.Println("[PLUGINS] Install discovered: ", sysPath)
				fmt.Println("[PLUGINS] Plugin Loaded:", pipelinePlugin, pipe)
				return pipe
				break
			}
		}

	}
	return pipe
}
