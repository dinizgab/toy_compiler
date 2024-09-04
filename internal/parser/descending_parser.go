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
	switch p.Lookahead.Type {
	case token.TokenFn:
		return p.parseFunctionDefinition()
	case token.TokenIdent:
		return p.parseExpression()
	default:
		return fmt.Errorf("descendingParserImpl.Parse: Unexpected token: %s", p.Lookahead.Value)
	}
}

func (p *descendingParserImpl) parseFunctionDefinition() error {
	const entityName = "descendingParserImpl.parseFunctionDefinition"

	// TODO - put the error to return inside the match function
	if !p.Match(token.TokenFn) {
		return fmt.Errorf("%s: Expected 'fn', found: %s", entityName, p.Lookahead.Value)
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

func (p *descendingParserImpl) parseStatements() error {
    return nil
}

func (p *descendingParserImpl) parseAssign() error {
	const entityName = "descendingParserImpl.parseAssign"

	if !p.Match(token.TokenIdent) {
		return fmt.Errorf("%s: Expected identifier, found: %s", entityName, p.Lookahead.Value)
	}

	if !p.Match(token.TokenAssign) {
		return fmt.Errorf("%s: Expected '=', found: %s", entityName, p.Lookahead.Value)
	}

	err := p.parseExpression()
	if err != nil {
		return err
	}

	return nil
}

func (p *descendingParserImpl) parseExpression() error {
    if err := p.parseTerm(); err != nil {
        return err
    }

    for p.Lookahead.Type == token.TokenAdditionOperator || p.Lookahead.Type == token.TokenSubtractionOperator {
        p.Match(p.Lookahead.Type)

        if err := p.parseTerm(); err != nil {
            return err
        }
    }

	return nil
}

func (p *descendingParserImpl) parseTerm() error {
    if err := p.parseFactor(); err != nil { 
        return err
    }

    for p.Lookahead.Type == token.TokenMultiplicationOperator || p.Lookahead.Type == token.TokenDivisionOperator {
        p.Match(p.Lookahead.Type)
        if err := p.parseFactor(); err != nil {
            return err
        }
    }

    return nil
}

func (p *descendingParserImpl) parseFactor() error {
    const entityName = "descendingParserImpl.parseFactor"
    if p.Lookahead.Type == token.TokenIdent || p.Lookahead.Type == token.TokenNumber {
        p.Match(p.Lookahead.Type)

        return nil
    } else if p.Lookahead.Type == token.TokenOpenParen {
        p.Match(token.TokenOpenParen)
        p.parseExpression()
        p.Match(token.TokenCloseParen)
    } else {
        return fmt.Errorf("%s: Unexpected token type: %s", entityName, p.Lookahead.Value)
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
