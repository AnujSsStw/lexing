package lexer

type Lexer struct {
	input string
	currP int
	nextP int
	ch    byte
}

func (l *Lexer) NextToken() {}

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
