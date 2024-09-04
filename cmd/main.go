package main
import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dinizgab/toy_compiler/internal/lexer"
	"github.com/dinizgab/toy_compiler/internal/parser"
	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
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


	st := symboltable.New()
	parser := parser.NewDescendingParser(tokens, st)
    err = parser.Parse()
    if err != nil {
        log.Fatalf("Error parsing file: %v", err)
    }

    fmt.Println("Symbol Table:")
    for k, v := range st {
        fmt.Printf("%s: %v\n", k, v)
    }

}
