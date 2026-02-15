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
	EQUAL      TokenType = "=="
	NOT_EQ     TokenType = "!="
	MINUS      TokenType = "-"
	SLASH      TokenType = "/"
	TRUE       TokenType = "TRUE"
	FALSE      TokenType = "FALSE"
	IF         TokenType = "IF"
	ELSE       TokenType = "ELSE"
	ASTERISK   TokenType = "*"
	BANG       TokenType = "!"
	LT         TokenType = "<"
	GT         TokenType = ">"
	RETURN     TokenType = "RETURN"
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

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"else":   ELSE,
	"if":     IF,
	"true":   TRUE,
	"false":  FALSE,
}

func GetIdentifierType(word string) TokenType {
	v, ok := keywords[word]
	if !ok {
		return IDENTIFIER
	}
	return v
}

func NewToken(tType TokenType, literal byte) Token {
	return Token{
		Type:    tType,
		Literal: string(literal),
	}
}
