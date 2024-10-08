package lexer

import (
	"errors"
	"fmt"

	"github.com/dinizgab/toy_compiler/internal/token"
	"github.com/dinizgab/toy_compiler/internal/utils"
)

type Lexer struct {
	Input   string
	Cursor  int
	Forward int
}

func New(input []byte) *Lexer {
	return &Lexer{
		Input:   string(input),
		Cursor:  0,
		Forward: 1,
	}
}

func (l *Lexer) Lex() ([]*token.Token, error) {
	tokens := make([]*token.Token, 0)

	for l.Cursor < len(l.Input) {
		char := l.peek()

		switch {
		case utils.IsNum(char):
			token, err := l.readNum()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case utils.IsAlpha(char):
			token, err := l.readIdentifier()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case utils.IsOperator(char):
			token, err := l.readOperator()
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
		case utils.IsBracket(char):
			tokenType := utils.GetBracketType(char)

			token := token.NewToken(tokenType, string(char))
			tokens = append(tokens, token)
		case char == 10 || char == 32:
			// ignoring new lines and spaces
		default:
			return nil, fmt.Errorf("lexer.Lex: Invalid character: %s", string(char))
		}

		l.nextLexeme()
	}

	return tokens, nil
}

func (l *Lexer) peek() byte {
	return l.Input[l.Cursor]
}

func (l *Lexer) peekNextChar() byte {
	if l.Forward >= len(l.Input) {
		return 0
	}

	return l.Input[l.Forward]
}

func (l *Lexer) advance() {
	l.Forward++
}

func (l *Lexer) nextLexeme() {
	l.Cursor = l.Forward
	l.Forward += 1
}

func (l *Lexer) readIdentifier() (*token.Token, error) {
	if !utils.IsAlpha(l.peek()) {
		return nil, fmt.Errorf("lexer.readIdentifier: Invalid lexeme, a identifier must start with a letter: %d", l.Cursor)
	}

	for l.Forward < len(l.Input) && (utils.IsAlpha(l.peekNextChar()) || utils.IsNum(l.peekNextChar())) {
		l.advance()
	}

	lexeme := l.Input[l.Cursor:l.Forward]
	// If lexeme is a keyword
	// TODO - Create a function that create a new keyword token (now i have duplicated code in these params)
	if value, ok := token.IsKeyword(lexeme); ok {
		return token.NewToken(value, value), nil
	}

	return token.NewToken(token.TokenIdent, lexeme), nil
}

func (l *Lexer) readNum() (*token.Token, error) {
	seenDot := false

	for l.Forward < len(l.Input) {
		if l.peekNextChar() == '.' {
			if seenDot {
				return nil, errors.New("lexer.readNum: Number with more than one '.'")
			}

			seenDot = true
		}

		if utils.IsAlpha(l.peekNextChar()) {
			return nil, fmt.Errorf("lexer.readNum: Invalid lexeme, a number cannot be followed with a letter: %d", l.Cursor)
		}

		if !utils.IsNum(l.peekNextChar()) && l.peekNextChar() != '.' {
			break
		}

		l.advance()
	}

	lexeme := l.Input[l.Cursor:l.Forward]

	token := token.NewToken(token.TokenNumber, lexeme)
	return token, nil
}

func (l *Lexer) readOperator() (*token.Token, error) {
	if utils.IsOperator(l.peekNextChar()) {
		l.advance()
	}

	var tokenType string
	lexeme := l.Input[l.Cursor:l.Forward]

	switch lexeme {
	case "==", "!=", "&&", "||", ">", "<", "<=", ">=", "!":
		tokenType = token.TokenLogicalOperator
	case "=":
		tokenType = token.TokenAssign
    case "+":
        tokenType = token.TokenAdditionOperator
    case "-":
        tokenType = token.TokenSubtractionOperator
    case "*":
        tokenType = token.TokenMultiplicationOperator
    case "/":
        tokenType = token.TokenDivisionOperator
	default:
        return nil, fmt.Errorf("lexer.readOperator: Invalid operator: %s", lexeme)
	}

	token := token.NewToken(tokenType, lexeme)
	return token, nil
}
