package parser

import (
	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
	"github.com/dinizgab/toy_compiler/internal/token"
)

type Parser interface {
	ParseAscending() error
	ParseDescending() error
}

type parserImpl struct {
	Tokens      []*token.Token
	Lookahead   *token.Token
	symbolTable symboltable.SymbolTable
}

func New(tokens []*token.Token, st symboltable.SymbolTable) Parser {
	return &parserImpl{
		Tokens:      tokens,
		Lookahead:   tokens[0],
		symbolTable: st,
	}
}

func (p *parserImpl) ParseAscending() error {
	return nil
}

func (p *parserImpl) ParseDescending() error {
	return nil
}

func (p *parserImpl) match(t *token.Token) bool {
	return p.Lookahead != nil && p.Lookahead.Type == t.Type && p.Lookahead.Value == t.Value
}