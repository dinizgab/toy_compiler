package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dinizgab/toy_compiler/internal/lexer"
)

var filename string

func main() {
	flag.StringVar(&filename, "filename", "test.toy", "The path to the file to be compiled")
	flag.Parse()

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lex := lexer.New(fileContent)
	tokens, err := lex.Lex()
	if err != nil {
		log.Fatalf("Error lexing file: %v", err)
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}
