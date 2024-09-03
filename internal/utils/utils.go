package utils

import "github.com/dinizgab/toy_compiler/internal/token"

func IsAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func IsNum(char byte) bool {
	return (char >= '0' && char <= '9')
}

func IsOperator(char byte) bool {
	operators := []byte{'=', '<', '>', '+', '-', '/', '*', '!', '|', '&'}

	for _, value := range operators {
		if value == char {
			return true
		}
	}

	return false
}

func IsBracket(char byte) bool {
	brackets := []byte{'(', ')', '{', '}', '[', ']'}

	for _, value := range brackets {
		if value == char {
			return true
		}
	}

	return false
}

func GetBracketType(char byte) string {
    switch char {
    case '(':
        return token.TokenOpenParen
    case ')':
        return token.TokenCloseParen
    case '{':
        return token.TokenOpenBrack
    case '}':
        return token.TokenCloseBrack
    case '[':
        return "OPEN_SQUARE_BRACK"
    case ']':
        return "CLOSE_SQUARE_BRACK"
    }

    return ""
}
