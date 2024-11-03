package lexer

import "monkey/token"

// Lexer is a struct representing a lexical analyzer that processes an input string
// and breaks it down into tokens for easier parsing and interpretation.
// It keeps track of the current position in the input, the reading position (next character to read),
// and the current character under examination.
type Lexer struct {
	input        string // the input string to be tokenized
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position (after current char)
	ch           byte   // current char under examination
}

// New initializes a new Lexer instance with the given input string.
// It calls readChar to set the first character and returns the Lexer instance.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // initialize the first character
	return l
}

// NextToken examines the current character in the input string
// and returns the next token based on the character type (identifier, digit, etc.).
// It also skips over whitespace and returns an ILLEGAL token for unrecognized characters.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// Skip any whitespace characters
	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	default:
		// Check if the character is the start of an identifier (e.g., a variable name)
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// If the character is a digit, read the full number
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else if l.ch == 0 {
			// if it is end of line
			tok.Type = token.EOF
			tok.Literal = ""
		} else {
			// Return an ILLEGAL token for unrecognized characters
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar() // Move to the next character for the next tokenization call
	return tok
}

// newToken creates a new token of the given type with the literal value as the character.
// This is used for single-character tokens.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier reads a sequence of letters (a valid identifier) and returns it as a string.
// This function stops reading when it encounters a non-letter character.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readChar updates the Lexer's current character by advancing the readPosition.
// If the end of the input is reached, it sets the current character to 0.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for NUL, indicates end of input
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// isLetter checks if the given character is a letter (a-z, A-Z) or an underscore (_),
// which are valid starting characters for identifiers.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// skipWhitespace advances the position until it encounters a non-whitespace character.
// It skips spaces, tabs, newlines, and carriage returns.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isDigit checks if the given character is a digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readNumber reads a sequence of digit characters and returns it as a string.
// It stops reading when it encounters a non-digit character.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}