package lexer

import (
	"fmt"
	"testing"

	"github.com/dinizgab/toy_compiler/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestReadIdentifier(t *testing.T) {
	l := New([]byte("abc"))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenIdent, tokens[0].Type)
	assert.Equal(t, "abc", tokens[0].Value)
}

func TestReadNum(t *testing.T) {
	l := New([]byte("123"))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenNumber, tokens[0].Type)
	assert.Equal(t, "123", tokens[0].Value)
}


func TestReadNumDoubleDot(t *testing.T) {
	l := New([]byte("1.23.0"))
	tokens, err := l.Lex()

	assert.Error(t, err)
	assert.Equal(t, "lexer.readNum: Number with more than one '.'", err.Error())
	assert.Nil(t, tokens)
}


func TestReadNumWithAlpha(t *testing.T) {
	l := New([]byte("1.23a"))
	tokens, err := l.Lex()

	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("lexer.readNum: Invalid lexeme, a number cannot be followed with a letter: %d", l.Cursor), err.Error())
	assert.Nil(t, tokens)
}

func TestReadOperatorArithmetic(t *testing.T) {
	l := New([]byte("+"))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenArithmeticOperator, tokens[0].Type)
	assert.Equal(t, "+", tokens[0].Value)
}

func TestReadOperatorLogical(t *testing.T) {
	l := New([]byte("<"))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenLogicalOperator, tokens[0].Type)
	assert.Equal(t, "<", tokens[0].Value)
}

func TestReadOperatorAssign(t *testing.T) {
	l := New([]byte("="))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenAssign, tokens[0].Type)
	assert.Equal(t, "=", tokens[0].Value)
}

func TestReadOperatorDoubleChar(t *testing.T) {
	l := New([]byte("=="))
	tokens, err := l.Lex()

	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, token.TokenLogicalOperator, tokens[0].Type)
	assert.Equal(t, "==", tokens[0].Value)
}