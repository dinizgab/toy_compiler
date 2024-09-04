package parser

import (
)

type Parser interface {
	Parse() error
    Match(t string) error 
}
