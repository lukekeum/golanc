package compiler

import (
	"fmt"
	"os"

	"github.com/lukekeum/golanc/compiler/parser"
	"github.com/lukekeum/golanc/token"
)

func Execute(filename string) int {

	fmt.Println("Compiler: start compile")

	data, err := os.ReadFile(filename)

	if err != nil {
		panic("Error occured while reading file")
	}

	content := string(data[:])

	var tokenStore []token.Token = []token.Token{}

	// Lex Analyzer

	Lexical(content, tokenStore)

	// Parser
	parser := parser.New(tokenStore)

	parser.Init()

	fmt.Println("Compiler: compile complete")

	return 0
}
