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
	if err := p.Match(token.TokenFn); err != nil {
		return err
	}

	if err := p.Match(token.TokenIdent); err != nil {
		return err
	}

	if err := p.Match(token.TokenOpenParen); err != nil {
		return err
	}

	if err := p.parseFunctionParameters(); err != nil {
		return err
	}

	if err := p.Match(token.TokenCloseParen); err != nil {
		return err
	}

	if err := p.Match(token.TokenOpenBrack); err != nil {
		return err
	}

	if err := p.parseStatements(); err != nil {
		return err
	}

	if err := p.Match(token.TokenCloseBrack); err != nil {
		return err
	}

	return nil
}

func (p *descendingParserImpl) parseFunctionParameters() error {
	if p.peek().Value == token.TokenCloseParen {
		return nil
	}

	if err := p.Match(token.TokenIdent); err != nil {
		return err
	}

	for p.peek().Value != token.TokenCloseParen {
		if err := p.Match(token.TokenColon); err != nil {
			return err
		}

		if err := p.Match(token.TokenIdent); err != nil {
			return err
		}
	}
	return nil
}

func (p *descendingParserImpl) parseStatements() error {
	return nil
}

func (p *descendingParserImpl) parseAssign() error {
	if err := p.Match(token.TokenIdent); err != nil {
		return err
	}

	if err := p.Match(token.TokenAssign); err != nil {
		return err
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
		if err := p.Match(p.Lookahead.Type); err != nil {
			return err
		}

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
		if err := p.Match(p.Lookahead.Type); err != nil {
			return err
		}

		if err := p.parseFactor(); err != nil {
			return err
		}
	}

	return nil
}

func (p *descendingParserImpl) parseFactor() error {
	const entityName = "descendingParserImpl.parseFactor"
	if p.Lookahead.Type == token.TokenIdent || p.Lookahead.Type == token.TokenNumber {
		err := p.Match(p.Lookahead.Type)
		if err != nil {
			return err
		}

		return nil
	} else if p.Lookahead.Type == token.TokenOpenParen {
		err := p.Match(token.TokenOpenParen)
		if err != nil {
			return err
		}

		err = p.parseExpression()
		if err != nil {
			return err
		}

		err = p.Match(token.TokenCloseParen)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("%s: Unexpected token type: %s", entityName, p.Lookahead.Value)
	}

	return nil
}

func (p *descendingParserImpl) Match(t string) error {
	const entityName = "descendingParserImpl.Match"

	if p.Lookahead == nil {
		return fmt.Errorf("%s: Unexpected EOF", entityName)
	}

	if p.Lookahead != nil && p.Lookahead.Type == t {
		// TODO - Check if this entry already exists in the symbolTable
		if p.Lookahead.Type == token.TokenIdent {
			p.handleSymbolTableAddition()
		}

		p.NextPosition()
		return nil
	}

	tokenName := token.LiteralNameFromType(t)
	return fmt.Errorf("%s: Expected %s, found: %s", entityName, tokenName, p.Lookahead.Value)
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

func (p *descendingParserImpl) peek() *token.Token {
	if p.Position+1 > len(p.Tokens) {
		return nil
	}

	return p.Tokens[p.Position+1]
}
