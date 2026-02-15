package token

type TokenType string

const (
	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	RBRACE     TokenType = "}"
	LBRACE     TokenType = "{"
	COMMA      TokenType = ","
	SEMICOLON  TokenType = ";"
	PLUS       TokenType = "+"
	ASSIGN     TokenType = "="
	FUNCTION   TokenType = "FUNCTION"
	IDENTIFIER TokenType = "IDENTIFIER"
	LET        TokenType = "LET"
	INT        TokenType = "INT"
	EOF        TokenType = "EOF"     // tells the parser to stop
	ILLEGAL    TokenType = "ILLEGAL" // unknown token type
)

type SourceLocation struct {
	Filename string
	Line     int
	Column   int
}

type Token struct {
	Type     TokenType
	Literal  string
	Location SourceLocation
}

func NewToken(tType TokenType, literal byte, filename string, line int, column int) Token {
	return Token{
		Type:    tType,
		Literal: string(literal),
		Location: SourceLocation{
			Column:   column,
			Line:     line,
			Filename: filename,
		},
	}
}
