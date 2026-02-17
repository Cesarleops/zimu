package lexer

import (
	"log"

	"github.com/cesarleops/zimu/internal/token"
)

var logger = log.Default()

type LexingLoc struct {
	line     int
	column   int
	filename *string
}
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
	loc          LexingLoc
}

func NewLexer(input string, filename string) *Lexer {
	l := &Lexer{
		position:     0,
		readPosition: 0,
		loc: LexingLoc{
			filename: &filename,
			line:     1,
			column:   1,
		},
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// logger.Print("overflow")
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// end of file reached
	if l.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Location.Line = l.loc.line
		tok.Location.Filename = *l.loc.filename
		tok.Location.Column = l.loc.column
		return tok
	}

	switch l.ch {
	case '=':
		next := l.peak()
		if next == '=' {
			tok = l.genTwoCharToken(tok, token.EQUAL, l.ch, next)
			// logger.Printf("DOUBLE TOKEN %v \n", tok)
			return tok
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case '!':
		next := l.peak()
		if next == '=' {
			tok = l.genTwoCharToken(tok, token.NOT_EQ, l.ch, next)
			return tok
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)

	default:
		if isLetter(l.ch) {
			tok.Location.Column = l.loc.column // We set the column at the start of the word
			tok.Literal = l.readIdentifierLiteral()
			tok.Type = token.GetIdentifierType(tok.Literal)
			tok.Location.Line = l.loc.line
			tok.Location.Filename = *l.loc.filename
			// fmt.Printf("WORD TOKEN %v \n", tok)
			return tok
		} else if isDigit(l.ch) {
			tok.Location.Column = l.loc.column // We set the column at the start of the number
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			tok.Location.Line = l.loc.line
			tok.Location.Filename = *l.loc.filename
			// fmt.Printf("DIGIT TOKEN %v \n", tok)
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}
	tok.Location.Filename = *l.loc.filename
	tok.Location.Column = l.loc.column
	tok.Location.Line = l.loc.line
	l.readChar()
	l.loc.column += 1
	// fmt.Printf("NORMAL TOKEN %v \n", tok)
	return tok

}

func isLetter(ch byte) bool {
	// go transforms the runes to their byte value when comparing
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (l *Lexer) readIdentifierLiteral() string {
	curr := l.position
	for isLetter(l.ch) {
		l.loc.column += 1
		l.readChar()
	}
	return l.input[curr:l.position]
}

func (l *Lexer) readNumber() string {
	curr := l.position

	for isDigit(l.ch) {
		l.loc.column += 1
		l.readChar()
	}

	return l.input[curr:l.position]

}

func (l *Lexer) genTwoCharToken(tok token.Token, tType token.TokenType, fChar byte, sChar byte) token.Token {
	tok.Literal = string(fChar) + string(sChar)
	// logger.Printf("literal %q\n", tok.Literal)
	tok.Type = tType
	tok.Location.Filename = *l.loc.filename
	tok.Location.Column = l.loc.column
	tok.Location.Line = l.loc.line
	l.loc.column += 2
	// Move twice forward since we read the next token with the peak function
	l.readChar()
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {

	// TODO: CHECK HOW MANY COLUMNS IS A TAB OR CARRIAGE RETURN
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		if l.ch == '\n' {
			l.loc.line += 1
			l.loc.column = 1
		} else {
			l.loc.column += 1
		}
		l.readChar()
	}
}

func (l *Lexer) peak() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
