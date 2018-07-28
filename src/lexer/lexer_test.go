package lexer

import (
	"github.com/samwson/ghost/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `five = 5
ten = 10

def add(x, y)
	return x + y
end

result = add(five, ten)
!-/*5
5 < 10 > 5

if (5 < 10) 
	return true
else
	return false
end

10 == 10
10 != 9
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.NEWLINE, "\n"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.METHOD, "def"},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.NEWLINE, "\n"},
		{token.END, "end"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.NEWLINE, "\n"},
		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.GREATERTHAN, ">"},
		{token.INTEGER, "5"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.NEWLINE, "\n"},
		{token.ELSE, "else"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.NEWLINE, "\n"},
		{token.END, "end"},
		{token.NEWLINE, "\n"},
		{token.NEWLINE, "\n"},
		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.NEWLINE, "\n"},
		{token.INTEGER, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INTEGER, "9"},
		{token.NEWLINE, "\n"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		token := lexer.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("Tests[%d] - token type wrong. Expected=%q, got=%q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("Tests[%d] - literal wrong. Expected=%q, got=%q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
