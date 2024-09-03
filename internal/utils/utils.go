package utils

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