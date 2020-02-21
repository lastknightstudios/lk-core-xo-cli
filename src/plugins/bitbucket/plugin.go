package main

import "fmt"

type repository string

func (g repository) Repo() {
	fmt.Println("Hello bitbucket")
}

// Repository Type Symbol
var Repository repository
