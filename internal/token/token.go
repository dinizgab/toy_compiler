package token

import (
	"fmt"
)

const (
	TokenIf                 = "IF"
	TokenElseIf             = "ELSE_IF"
	TokenElse               = "ELSE"
	TokenFn                 = "FN"
	TokenReturn             = "RETURN"
	TokenOpenParen          = "OPEN_PAREN"
	TokenCloseParen         = "CLOSE_PAREN"
	TokenOpenBrack          = "OPEN_BRACK"
	TokenCloseBrack         = "CLOSE_BRACK"
	TokenIdent              = "IDENT"
	TokenNumber             = "NUMBER"
	TokenLogicalOperator    = "LOG_OPERATOR"
    TokenAdditionOperator   = "ADD_OPERATOR"
    TokenSubtractionOperator = "SUB_OPERATOR"
    TokenMultiplicationOperator = "MUL_OPERATOR"
    TokenDivisionOperator   = "DIV_OPERATOR"
	TokenAssign             = "ASSIGN"
	TokenEOF                = "EOF"
)

var tokens = map[string]string{
	"if":      TokenIf,
	"else if": TokenElseIf,
	"else":    TokenElse,
	"fn":      TokenFn,
	"return":  TokenReturn,
	"(":       TokenOpenParen,
	")":       TokenCloseParen,
	"{":       TokenOpenBrack,
	"}":       TokenCloseBrack,
}

type Token struct {
	Type  string
	Value string
}

func NewToken(tokenType, value string) (*Token) {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}

func (t *Token) String() string {
	return fmt.Sprintf("<%s : %s>", t.Type, t.Value)
}

func IsKeyword(lexeme string) (string, bool) {
	value, ok := tokens[lexeme]

	return value, ok
}

