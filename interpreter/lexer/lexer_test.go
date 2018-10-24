package lexer

import (
	"testing"

	"ape/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
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
		"foobar"
		"foo bar"
		[1, 2];
		{"foo": "bar"}
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.PARENL, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.PARENR, ")"},
		{token.BRACEL, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.BRACER, "}"},
		{token.SEMICOLON, ";"},

		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.PARENL, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.PARENR, ")"},
		{token.SEMICOLON, ";"},

		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.IF, "if"},
		{token.PARENL, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.PARENR, ")"},
		{token.BRACEL, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.BRACER, "}"},
		{token.ELSE, "else"},
		{token.BRACEL, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.BRACER, "}"},

		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},

		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},

		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},

		{token.BRACKETL, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.BRACKETR, "]"},
		{token.SEMICOLON, ";"},

		{token.BRACEL, "{"},
		{token.STRING, "foo"},
		{token.COLON, ":"},
		{token.STRING, "bar"},
		{token.BRACER, "}"},
		{token.EOF, ""},

		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - Bad Token Type. Expected '%q', got '%q' instead",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Bad Literal. Expected '%q', got '%q' instead",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
