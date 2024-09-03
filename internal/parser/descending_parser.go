package parser

import (
	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
	"github.com/dinizgab/toy_compiler/internal/token"
)

type descendingParserImpl struct {
	Tokens      []*token.Token
	Lookahead   *token.Token
	symbolTable symboltable.SymbolTable
}

func NewDescendingParser(tokens []*token.Token, st symboltable.SymbolTable) Parser {
	return &descendingParserImpl{
		Tokens:      tokens,
		Lookahead:   tokens[0],
		symbolTable: st,
	}
}

func (p *descendingParserImpl) Parse() error {
    return nil
}

func (p *descendingParserImpl) Match(t *token.Token) bool {
	return p.Lookahead != nil && p.Lookahead.Type == t.Type && p.Lookahead.Value == t.Value
}
