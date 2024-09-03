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
    value, ok := IsKeyword("fn")

	assert.True(t, ok)
    assert.Equal(t, TokenFn, value)
}

func TestIsKeywordInvalid(t *testing.T) {
    value, ok := IsKeyword("test")

    assert.False(t, ok)
    assert.Equal(t, "", value)
}
