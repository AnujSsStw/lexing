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
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"

	//Identifiers + literals
	INT   = "INT"
	IDENT = "IDENT"

	// Operators
	MUX       = "*"
	EQ        = "=="
	NOT_EQ    = "!="
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	DECREMENT = "--"
	INCREMENT = "++"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func IdentifierType(ident string) TokenType {
	if val, ok := keywords[ident]; ok {
		return val
	}
	return IDENT
}
