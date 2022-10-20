package main

import (
	"os"

	"github.com/lukekeum/golanc/compiler"
)

func main() {
	fileName := os.Args

	if len(os.Args) <= 1 {
		panic("file name argument required")
	}

	compiler.Execute(fileName[1])
}
