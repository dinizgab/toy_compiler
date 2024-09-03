package parser

import (
	"fmt"

	symboltable "github.com/dinizgab/toy_compiler/internal/symbol_table"
	"github.com/dinizgab/toy_compiler/internal/token"
)

type descendingParserImpl struct {
	Tokens      []*token.Token
	Lookahead   *token.Token
	Position    int
	symbolTable symboltable.SymbolTable
}

func NewDescendingParser(tokens []*token.Token, st symboltable.SymbolTable) Parser {
	return &descendingParserImpl{
		Tokens:      tokens,
		Lookahead:   tokens[0],
		Position:    0,
		symbolTable: st,
	}
}

func (p *descendingParserImpl) NextPosition() {
	p.Position++
	if p.Position >= len(p.Tokens) {
		p.Lookahead = &token.Token{Type: token.TokenEOF, Value: token.TokenEOF}
	} else {
		p.Lookahead = p.Tokens[p.Position]
	}
}

func (p *descendingParserImpl) Parse() error {
	return p.parseFunctionDefinition()
}

func (p *descendingParserImpl) parseFunctionDefinition() error {
	const entityName = "descendingParserImpl.parseFunctionDefinition"

    // TODO - put the error to return inside the match function
	if !p.Match(token.TokenFn) {
	}

	if !p.Match(token.TokenIdent) {
		return fmt.Errorf("%s: Expected function name, found: %s", entityName, p.Lookahead.Value)
	}

	if !p.Match(token.TokenOpenParen) {
		return fmt.Errorf("%s: Expected '(', found: %s", entityName, p.Lookahead.Value)
	}

	if !p.Match(token.TokenCloseParen) {
		return fmt.Errorf("%s: Expected ')', found: %s", entityName, p.Lookahead.Value)
	}

	if !p.Match(token.TokenOpenBrack) {
		return fmt.Errorf("%s: Expected '{', found: %s", entityName, p.Lookahead.Value)
	}

	// TODO - Parse function body

	if !p.Match(token.TokenCloseBrack) {
		return fmt.Errorf("%s: Expected '}', found: %s", entityName, p.Lookahead.Value)
	}

	return nil
}

func (p *descendingParserImpl) Match(t string) bool {
	if p.Lookahead != nil && p.Lookahead.Type == t {
		if p.Lookahead.Type == token.TokenIdent {
			p.handleSymbolTableAddition()
		}

		p.NextPosition()
		return true
	}

	return false
}

func (p *descendingParserImpl) handleSymbolTableAddition() {
	functionName := p.Lookahead.Value
	p.symbolTable.AddEntry(
		functionName,
		symboltable.SymbolInformation{
			Name: functionName,
			Addr: 0,
		},
	)
}
