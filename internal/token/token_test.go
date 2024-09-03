package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token, err := NewToken(TokenIdent, "test")

	assert.Nil(t, err)
	assert.Equal(t, token.Type, TokenIdent)
	assert.Equal(t, token.Value, "test")
}

func TestIsKeywordValid(t *testing.T) {
	assert.True(t, IsKeyword("fn"))
	assert.True(t, IsKeyword("if"))
	assert.True(t, IsKeyword("else"))
	assert.True(t, IsKeyword("return"))
}

func TestIsKeywordInvalid(t *testing.T) {
	assert.False(t, IsKeyword("test"))
	assert.False(t, IsKeyword("fi"))
	assert.False(t, IsKeyword("esle"))
	assert.False(t, IsKeyword("foo"))
}