package main

import "fmt"

func help() {
	fmt.Println("\n", "USAGE: xo <command> [args] [repo] [pipeline]")
	fmt.Println("\n", "EXAMPLE: xo create project-name github buildkite")
	fmt.Println("\n", "available commands:")
	fmt.Println("\t", "create", "\t", "Create a new project")
	fmt.Println("\t", "version", "\t", "Shows xo cli tool version")
	fmt.Println("\t", "help", "\t\t", "Shows help")
	fmt.Println("\n", "supported repos: bitbucket, github")
	fmt.Println("\n", "supported pipeline: buildkite")
	fmt.Println("\n", "source: https://github.com/lastknightstudios/lk-core-xo-cli")
	fmt.Println()
}
