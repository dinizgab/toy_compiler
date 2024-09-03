package parser

import (
	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
	"github.com/dinizgab/toy_compiler/internal/token"
)

type ascendingParserImpl struct {
	Tokens      []*token.Token
	Lookahead   *token.Token
	symbolTable symboltable.SymbolTable
}

func NewAscendingParser(tokens []*token.Token, st symboltable.SymbolTable) Parser {
	return &ascendingParserImpl{
		Tokens:      tokens,
		Lookahead:   tokens[0],
		symbolTable: st,
	}
}

func (p *ascendingParserImpl) Parse() error {
    return nil
}

func (p *ascendingParserImpl) Match(t string) bool {
	return p.Lookahead != nil && p.Lookahead.Type == t
}
