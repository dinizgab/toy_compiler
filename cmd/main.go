package main

import (
	"fmt"

	"github.com/dinizgab/toy_compiler/internal/lexer"
)

const (
	testString = "fn foo() {\n if (test <= 1 || 4 <= 3) {\na = 3\nreturn 1.2\n} else {\n return 2\n}\n }\n fn foo342() { a n 3 }"
	testString2 = "()"
)

func main() {
	lex := lexer.Lexer{Input: testString, Cursor: 0, Forward: 1}
	tokens, err := lex.Lex()
	if err != nil {
		fmt.Println(err)
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}
