package token

type TokenType string

type Token struct {
	Type         TokenType
	LiteralValue string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	//Identifiers + literals
	INT   = "INT"
	IDENT = "IDENT"

	// Operators
	MUX      = "*"
	EQUALITY = "=="
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)
