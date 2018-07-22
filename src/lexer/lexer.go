// Package lexer does lexical analysis on an input string to produce tokens.

package lexer

import (
	"github.com/samwson/ghost/src/token"
)

type Lexer struct {
	input        string
	position     int  // The current position in input (points to current char)
	readPosition int  // The current reading position in input (after current char)
	ch           byte // Current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCharacter()
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.ch {
	case '=':
		tok = newToken(token.ASSIGN, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lexer.readCharacter()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
