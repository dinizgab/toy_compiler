package lexer

import "fmt"

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
    TokenArithmeticOperator = "ARITH_OPERATOR"
    TokenAssign             = "ASSIGN"
)

var tokens = map[string]string{
    "if":                 TokenIf,
    "else if":            TokenElseIf,
    "else":               TokenElse,
    "fn":                 TokenFn,
    "return":             TokenReturn,
    "(":                  TokenOpenParen,
    ")":                  TokenCloseParen,
    "{":                  TokenOpenBrack,
    "}":                  TokenCloseBrack,
    "id":                 TokenIdent,
    "number":             TokenNumber,
    "logicalOperator":    TokenLogicalOperator,
    "arithmeticOperator": TokenArithmeticOperator,
    "assign":             TokenAssign,
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