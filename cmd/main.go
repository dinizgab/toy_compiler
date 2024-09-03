package main

import (
	"errors"
	"fmt"

	"github.com/dinizgab/toy_compiler/internal/lexer"
)

const (
	testString = "fn foo() {\n if (test <= 1 || 4 <= 3) {\na = 3\nreturn 1.2\n} else {\n return 2\n}\n }\n fn foo342() { a n 3 }"
)





type Lexer struct {
	Input   string
	Cursor  int
	Forward int
}

func (l *Lexer) Lex() ([]*lexer.Token, error) {
	tokens := make([]*lexer.Token, 0)

	for l.Cursor < l.Forward && len(l.Input) > l.Cursor {
		char := l.Input[l.Cursor]

		if isAlpha(char) {
			tokens = append(tokens, l.readIdentifier())
		} else if isNum(char) {
			token, err := l.readNum()
			if err != nil {
				return nil, err
			}

			tokens = append(tokens, token)
		} else if isOperator(char) {
			tokens = append(tokens, l.readOperator())
		} else {
			// ignoring new lines and spaces
			if char != 10 && char != 32 {
				tokens = append(tokens, lexer.NewToken(string(char), string(char)))
			}
		}
		l.nextLexeme()
	}

	return tokens, nil
}

func (l *Lexer) nextLexeme() {
	l.Cursor = l.Forward
	l.Forward += 1
}

func (l *Lexer) nextChar() byte {
	return l.Input[l.Cursor+1]
}

func (l *Lexer) readIdentifier() *lexer.Token {
	for l.Forward < len(l.Input) && (isAlpha(l.Input[l.Forward]) || isNum(l.Input[l.Forward])) {
		l.Forward++
	}

	lexeme := l.Input[l.Cursor:l.Forward]
	// If lexeme is a keyword
	if tokenType, ok := lexer.tokens[lexeme]; ok {
		return lexer.NewToken(lexeme, tokenType)
	}

	return lexer.NewToken("id", lexeme)
}

func (l *Lexer) readNum() (*lexer.Token, error) {
	dotCounter := 0
	for l.Forward < len(l.Input) && isNum(l.Input[l.Forward]) || l.Input[l.Forward] == '.' {
		if l.Input[l.Forward] == '.' {
			dotCounter++

			if dotCounter > 1 {
				return nil, errors.New("lexer.readNum: Number with more than one '.'!")
			}
		}

		l.Forward++
	}

	lexeme := l.Input[l.Cursor:l.Forward]
	return lexer.NewToken("number", lexeme), nil
}

func (l *Lexer) readOperator() *lexer.Token {
	if isOperator(l.nextChar()) {
		l.Forward++
	}

	var tokenType string
	lexeme := l.Input[l.Cursor:l.Forward]

	switch lexeme {
	case "==", "!=", "&&", "||", ">", "<", "<=", ">=", "!":
		tokenType = "logicalOperator"
	case "=":
		tokenType = "assign"
	default:
		tokenType = "arithmeticOperator"
	}

	return lexer.NewToken(tokenType, lexeme)
}

func main() {
	lex := Lexer{Input: testString, Cursor: 0, Forward: 1}
	tokens, err := lex.Lex()
	if err != nil {
		fmt.Println(err)
	}

	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isNum(char byte) bool {
	return (char >= '0' && char <= '9')
}

func isOperator(char byte) bool {
	operators := []byte{'=', '<', '>', '+', '-', '/', '*', '!', '|', '&'}

	for _, value := range operators {
		if value == char {
			return true
		}
	}

	return false
}
