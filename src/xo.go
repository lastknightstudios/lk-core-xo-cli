package main

import (
	"fmt"
	"os"
)

func main() {

	// Parse Arguments
	args := os.Args

	if len(os.Args) < 2 {
		help(os.Args)
	} else {
		command := os.Args[1]
		switch command {
		case "version":
			version(args)
		case "config":
			config(args)
		case "create":
			create(args)
		case "help":
			help(args)
		default:
			fmt.Println("Invalid command")
		}
	}
}
