package main

import "fmt"

// FI: Automate supported plugins by dynamically creating list from so files in firectory

func help() {
	fmt.Println("\n", version())
	fmt.Println("\n", "USAGE: xo <command> [args] [repo] [pipeline]")
	fmt.Println("\n", "EXAMPLE: xo create project-name github buildkite")
	fmt.Println("\n", "available commands:")
	fmt.Println("\t", "create", "\t", "Create a new project")
	fmt.Println("\t", "help", "\t", "Shows command help")
	fmt.Println("\t", "version", "\t", "Shows xo cli tool version")
	fmt.Println("\n", "supported repos: github,")
	fmt.Println("\n", "supported pipeline: buildkite")
	fmt.Println("\n", "source: https://github.com/lastknightstudios/lk-core-xo-cli")
	fmt.Println()
}
