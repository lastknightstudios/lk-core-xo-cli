package main

import (
	"fmt"
	"os"
)

func createDir(dirName string) {
	fmt.Println("[OS] Creating directory", dirName)
	os.Mkdir(dirName, os.ModePerm)

}
