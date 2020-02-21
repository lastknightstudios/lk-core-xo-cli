package main

import (
	"fmt"
	"os"
)

var vcsPlugin = "github"

func main() {

	// Parse Arguments
	args := os.Args

	if len(os.Args) < 2 {
		help(args)
	} else {
		command := os.Args[1]
		switch command {
		case "version":
			version(args)
		case "help":
			help(args)
		//case "config":
		//	config(args)
		//case "create":
		//	create(args)
		default:
			fmt.Println("Invalid command")
		}
	}
}
