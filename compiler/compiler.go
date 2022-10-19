package compiler

import (
	"fmt"
	"os"
)

func Compile(filename string) int {

	fmt.Println("Compiler: start compile")

	data, err := os.ReadFile(filename)

	if err != nil {
		panic("Error occured while reading file")
	}

	fmt.Println(string(data)) // TODO: Add Compiler Section

	fmt.Println("Compiler: compile complete")

	return 0
}
