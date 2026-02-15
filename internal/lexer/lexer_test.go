package lexer

import (
	"os"
	"testing"

	"github.com/cesarleops/zimu/internal/token"
)

func TestNextToken(t *testing.T) {
	input, err := os.ReadFile("../../test/monkey_random_words.mon")
	if err != nil {
		os.Exit(1)
	}
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
			expectedLine:    1,
			expectedColumn:  0,
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
			expectedLine:    1,
			expectedColumn:  1,
		},
		{
			expectedType:    "",
			expectedLiteral: "",
			expectedLine:    0,
			expectedColumn:  0,
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
			expectedLine:    2,
			expectedColumn:  0,
		},
		{
			expectedType:    "",
			expectedLiteral: "",
			expectedLine:    0,
			expectedColumn:  0,
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
			expectedLine:    3,
			expectedColumn:  0,
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
			expectedLine:    3,
			expectedColumn:  1,
		},
	}

	lex := NewLexer(string(input), "test")

	for i, tt := range tests {
		tok := lex.NextToken()
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf("tests[%d]-token literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

		if tt.expectedType != tok.Type {
			t.Fatalf("tests[%d]-token type wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)

		}

		if tt.expectedLine != tok.Location.Line {
			t.Fatalf("tests[%d]-token line wrong, expected=%d, got=%d", i, tt.expectedLine, tok.Location.Line)

		}

		if tt.expectedColumn != tok.Location.Column {
			t.Fatalf("tests[%d]-token column wrong, expected=%d, got=%d", i, tt.expectedColumn, tok.Location.Column)

		}
	}

}
