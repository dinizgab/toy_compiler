package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token, err := NewToken("id", "test")

	assert.Nil(t, err)
	assert.Equal(t, token.Type, TokenIdent)
}

func TestNewTokenInvalidType(t *testing.T) {
	token, err := NewToken("invalid", "test")

	assert.Error(t, err, "invalid token type: invalid")
	assert.Nil(t, token)
}