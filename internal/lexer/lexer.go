package lexer

import (
	"log"

	"github.com/cesarleops/zimu/internal/token"
)

var logger = log.Default()

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	line         int
	filename     *string
	column       int
}

func NewLexer(input string, filename string) *Lexer {

	l := &Lexer{
		position:     0,
		readPosition: 0,
		filename:     &filename,
		input:        input,
		line:         1,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	if l.ch == '\n' {
		l.line += 1
		l.column = 0
		l.readChar()
		return tok
	}
	// end of file reached
	if l.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Location.Line = l.line
		tok.Location.Filename = *l.filename
		tok.Location.Column = l.column
	}

	switch l.ch {
	case '=':
		tok = token.NewToken(token.ASSIGN, l.ch, *l.filename, l.line, l.column)
	case '+':
		tok = token.NewToken(token.PLUS, l.ch, *l.filename, l.line, l.column)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch, *l.filename, l.line, l.column)

	case ')':
		tok = token.NewToken(token.RPAREN, l.ch, *l.filename, l.line, l.column)

	case '{':
		tok = token.NewToken(token.LBRACE, l.ch, *l.filename, l.line, l.column)

	case '}':
		tok = token.NewToken(token.RBRACE, l.ch, *l.filename, l.line, l.column)
	}
	l.readChar()
	l.column += 1
	return tok

}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
