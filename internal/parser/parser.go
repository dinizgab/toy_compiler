package parser

import (
	"github.com/dinizgab/toy_compiler/internal/token"
)

type Parser interface {
	Parse() error
    Match(t *token.Token) bool
}
