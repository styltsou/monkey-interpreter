package lexer

import "github.com/styltsou/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current character that is parsed
}

// Utility function to read the next character from the input
//* Note: this syntax after the func keyword is called function receiver
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 0 in ASCI corresponds to NUL
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// Returns the next character to be read without actually changing the lexer's cursor position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		// corresponds to EOF
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) eatWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

// If the parsed character is a letter
// Continue reading chars, until your reach a non-letter char
// return the chars read (literal)
func (l *Lexer) readIdentifier() string {
	startIdx := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[startIdx:l.position]
}

func (l *Lexer) readNumber() string {
	startIdx := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[startIdx:l.position]
}

// ! I dont like this very much. The first input parsing should be done when calling NextToken
// Initialize a Lexer and read the first input char
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// ! I dont like this implementation suggested by the book
// ! NextToken should parse the next char end return a token of it
// Thus doing l := New(input) and tok = l.NextToken(), tokenizes the first char of the input
func (l *Lexer) NextToken() token.Token {
	l.eatWhitespace()

	// l.readChar()
	var tok token.Token

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case 0:
		tok = newToken(token.EOF, 0)
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifierType(tok.Literal)
			// We exit early because we have already read the next char (past the end of the identifier)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return tok
}
