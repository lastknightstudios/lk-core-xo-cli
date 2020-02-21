package main

import "fmt"

func help(args []string) {
	print(args, "\n")
	fmt.Println("\n", "USAGE: xo <command> [args] e.g. xo version")
	fmt.Println("\n", "available commands:")
	//fmt.Println("\t", "config", "\t", "Configures xo cli options")
	fmt.Println("\t", "create", "\t", "Create a new project")
	fmt.Println("\t", "version", "\t", "Shows xo cli tool version")
	fmt.Println("\t", "help", "\t\t", "Shows help")
	fmt.Println()
}
