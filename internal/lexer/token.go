package lexer

import "fmt"

var tokens = map[string]string{
	"if":                 "IF",
	"else if":            "ELSE_IF",
	"else":               "ELSE",
	"fn":                 "FN",
	"return":             "RETURN",
	"(":                  "OPEN_PAREN",
	")":                  "CLOSE_PAREN",
	"{":                  "OPEN_BRACK",
	"}":                  "CLOSE_BRACK",
	"id":                 "IDENT",
	"number":             "NUMBER",
	"logicalOperator":    "LOG_OPERATOR",
	"arithmeticOperator": "ARITH_OPERATOR",
	"assign":             "ASSIGN",
}

type Token struct {
	Type  string
	Value string
}

func NewToken(tokenType, value string) *Token {
	tokenType = tokens[tokenType]

	return &Token{
		Type:  tokenType,
		Value: value,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("<%s : %s>", t.Type, t.Value)
}