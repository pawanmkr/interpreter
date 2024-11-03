package token

type TokenType string
type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	// DELIMITERS
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// if a word is ident or keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}