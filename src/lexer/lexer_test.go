package lexer

import (
	"github.com/samwson/ghost/src/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(),;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
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
