package parser

import (
	"fmt"

	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
	"github.com/dinizgab/toy_compiler/internal/token"
)

type ascendingParserImpl struct {
	Tokens      []*token.Token
	Lookahead   *token.Token
	symbolTable symboltable.SymbolTable
	Position    int
}

func NewAscendingParser(tokens []*token.Token, st symboltable.SymbolTable) Parser {
	return &ascendingParserImpl{
		Tokens:      tokens,
		Lookahead:   tokens[0],
		Position:    0,
		symbolTable: st,
	}
}

func (p *ascendingParserImpl) Parse() error {
	return nil
}

func (p *ascendingParserImpl) Match(t string) error {
	const entityName = "ascendingParserImpl.Match"

	if p.Lookahead != nil && p.Lookahead.Type == t {
		p.Lookahead = p.Tokens[p.Position+1]
		return nil
	}

	return fmt.Errorf("%s: Unexpected token: %s", entityName, p.Lookahead.Value)
}
