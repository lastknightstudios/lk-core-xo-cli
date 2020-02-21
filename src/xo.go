package main

import (
	"fmt"
	"os"
)

func main() {

	// Parse Arguments
	args := os.Args

	if len(os.Args) < 2 {
		help(args)
	} else {
		command := os.Args[1]
		load(args)
		switch command {
		case "version":
			version(args)
		case "help":
			help(args)
		//case "config":
		//	config(args)
		case "create":
			create(args)
		default:
			fmt.Println("Invalid command")
			fmt.Println("try xo help")
		}
	}
}
