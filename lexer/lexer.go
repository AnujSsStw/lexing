package lexer

import (
	"github.com/AnujSsStw/lexing/token"
)

type Lexer struct {
	input string
	currP int
	nextP int
	ch    byte
}

func (l *Lexer) whiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextP >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextP]
	}
}

// tokenTypes -> 0 for peekChar, 1 for normal
func (l *Lexer) makeTwoCharToken(peekchar byte, ch byte, tok *token.Token, tokeTypes []token.TokenType) {
	if l.peekChar() == peekchar {
		ch := l.ch
		l.readChar()
		*tok = token.Token{Type: tokeTypes[0], LiteralValue: string(ch) + string(l.ch)}
	} else {
		*tok = getToken(tokeTypes[1], l.ch)
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.whiteSpace()

	switch l.ch {
	case '=':
		l.makeTwoCharToken('=', l.ch, &tok, []token.TokenType{token.EQ, token.ASSIGN})
	case ';':
		tok = getToken(token.SEMICOLON, l.ch)
	case '(':
		tok = getToken(token.LPAREN, l.ch)
	case ')':
		tok = getToken(token.RPAREN, l.ch)
	case ',':
		tok = getToken(token.COMMA, l.ch)
	case '+':
		l.makeTwoCharToken('+', l.ch, &tok, []token.TokenType{token.INCREMENT, token.PLUS})
	case '{':
		tok = getToken(token.LBRACE, l.ch)
	case '}':
		tok = getToken(token.RBRACE, l.ch)
	case 0:
		tok.LiteralValue = ""
		tok.Type = token.EOF
	case '-':
		l.makeTwoCharToken('-', l.ch, &tok, []token.TokenType{token.DECREMENT, token.MINUS})
	case '!':
		l.makeTwoCharToken('=', l.ch, &tok, []token.TokenType{token.NOT_EQ, token.BANG})
	case '/':
		tok = getToken(token.SLASH, l.ch)
	case '*':
		tok = getToken(token.ASTERISK, l.ch)
	case '<':
		tok = getToken(token.LT, l.ch)
	case '>':
		tok = getToken(token.GT, l.ch)
	default:
		if isLetter(l.ch) {
			tok.LiteralValue = l.readIdentifier()
			tok.Type = token.IdentifierType(tok.LiteralValue)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.LiteralValue = l.readNum()
			return tok
		} else {
			tok = getToken(token.EOF, l.ch)
		}

	}
	l.readChar()
	return tok
}
func (l *Lexer) readNum() string {
	p := l.currP
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[p:l.currP]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	p := l.currP
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[p:l.currP]

}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readChar() {
	if l.nextP >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextP]
	}
	l.currP = l.nextP
	l.nextP++
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func getToken(tokenType token.TokenType, val byte) token.Token {
	return token.Token{Type: tokenType, LiteralValue: string(val)}
}
