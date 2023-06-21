package lexer

import "github.com/styltsou/monkey-interpreter/token"

// Utility function to create a new Token struct
func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

// Returns true if the provided byte represents a letter (_ is counted as a letter)
func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
