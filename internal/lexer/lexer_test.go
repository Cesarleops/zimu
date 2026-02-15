package lexer

import (
	"testing"

	"github.com/cesarleops/zimu/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		// Line 1: let five = 5;
		{token.LET, "let", 1, 1},
		{token.IDENTIFIER, "five", 1, 5},
		{token.ASSIGN, "=", 1, 10},
		{token.INT, "5", 1, 12},
		{token.SEMICOLON, ";", 1, 13},

		// Line 2: let ten = 10;
		{token.LET, "let", 2, 1},
		{token.IDENTIFIER, "ten", 2, 5},
		{token.ASSIGN, "=", 2, 9},
		{token.INT, "10", 2, 11},
		{token.SEMICOLON, ";", 2, 13},

		// Line 4: let add = fn(x, y) {
		{token.LET, "let", 4, 1},
		{token.IDENTIFIER, "add", 4, 5},
		{token.ASSIGN, "=", 4, 9},
		{token.FUNCTION, "fn", 4, 11},
		{token.LPAREN, "(", 4, 13},
		{token.IDENTIFIER, "x", 4, 14},
		{token.COMMA, ",", 4, 15},
		{token.IDENTIFIER, "y", 4, 17},
		{token.RPAREN, ")", 4, 18},
		{token.LBRACE, "{", 4, 20},

		// Line 5:   x + y; (2 spaces indent)
		{token.IDENTIFIER, "x", 5, 3},
		{token.PLUS, "+", 5, 5},
		{token.IDENTIFIER, "y", 5, 7},
		{token.SEMICOLON, ";", 5, 8},

		// Line 6: };
		{token.RBRACE, "}", 6, 1},
		{token.SEMICOLON, ";", 6, 2},

		// Line 8: let result = add(five, ten);
		{token.LET, "let", 8, 1},
		{token.IDENTIFIER, "result", 8, 5},
		{token.ASSIGN, "=", 8, 12},
		{token.IDENTIFIER, "add", 8, 14},
		{token.LPAREN, "(", 8, 17},
		{token.IDENTIFIER, "five", 8, 18},
		{token.COMMA, ",", 8, 22},
		{token.IDENTIFIER, "ten", 8, 24},
		{token.RPAREN, ")", 8, 27},
		{token.SEMICOLON, ";", 8, 28},

		// Line 9: !-/*5;
		{token.BANG, "!", 9, 1},
		{token.MINUS, "-", 9, 2},
		{token.SLASH, "/", 9, 3},
		{token.ASTERISK, "*", 9, 4},
		{token.INT, "5", 9, 5},
		{token.SEMICOLON, ";", 9, 6},

		// Line 10: 5 < 10 > 5;
		{token.INT, "5", 10, 1},
		{token.LT, "<", 10, 3},
		{token.INT, "10", 10, 5},
		{token.GT, ">", 10, 8},
		{token.INT, "5", 10, 10},
		{token.SEMICOLON, ";", 10, 11},

		// Line 12: if (5 < 10) {
		{token.IF, "if", 12, 1},
		{token.LPAREN, "(", 12, 4},
		{token.INT, "5", 12, 5},
		{token.LT, "<", 12, 7},
		{token.INT, "10", 12, 9},
		{token.RPAREN, ")", 12, 11},
		{token.LBRACE, "{", 12, 13},

		// Line 13: 	return true; (1 tab indent)
		{token.RETURN, "return", 13, 2},
		{token.TRUE, "true", 13, 9},
		{token.SEMICOLON, ";", 13, 13},

		// Line 14: } else {
		{token.RBRACE, "}", 14, 1},
		{token.ELSE, "else", 14, 3},
		{token.LBRACE, "{", 14, 8},

		// Line 15: 	return false; (1 tab indent)
		{token.RETURN, "return", 15, 2},
		{token.FALSE, "false", 15, 9},
		{token.SEMICOLON, ";", 15, 14},

		// Line 16: }
		{token.RBRACE, "}", 16, 1},

		// Line 18: 10 == 10;
		{token.INT, "10", 18, 1},
		{token.EQUAL, "==", 18, 4},
		{token.INT, "10", 18, 7},
		{token.SEMICOLON, ";", 18, 9},

		// Line 19: 10 != 9;
		{token.INT, "10", 19, 1},
		{token.NOT_EQ, "!=", 19, 4},
		{token.INT, "9", 19, 7},
		{token.SEMICOLON, ";", 19, 8},

		// EOF
		{token.EOF, "", 20, 1},
	}
	lex := NewLexer(input, "test")

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
