package main

import (
	"fmt"
	"os"
)

func createDir(dirName string) {
	fmt.Println("[OS] Creating directory", dirName)
	os.Mkdir(dirName, os.ModePerm)
}

func getCurrentDir() string {
	path, err := os.Getwd()
	if err != nil {
		println(err)
		os.Exit(1)
	}
	return path
}
