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

func (lexer *Lexer) peekCharacter() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		if lexer.peekCharacter() == '=' {
			ch := lexer.ch
			lexer.readCharacter()
			literal := string(ch) + string(lexer.ch)
			tok = token.Token{Type: token.EQUAL, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, lexer.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tok = newToken(token.LESSTHAN, lexer.ch)
	case '>':
		tok = newToken(token.GREATERTHAN, lexer.ch)
	case '!':
		if lexer.peekCharacter() == '=' {
			ch := lexer.ch
			lexer.readCharacter()
			literal := string(ch) + string(lexer.ch)
			tok = token.Token{Type: token.NOT_EQUAL, Literal: literal}
		} else {
			tok = newToken(token.BANG, lexer.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '\n':
		tok = newToken(token.NEWLINE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexer.ch) {
			tok.Type = token.INTEGER
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.ch)
		}
	}

	lexer.readCharacter()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readIdentifier() string {
	start_position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readCharacter()
	}

	return lexer.input[start_position:lexer.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' {
		lexer.readCharacter()
	}
}

func (lexer *Lexer) readNumber() string {
	start_position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readCharacter()
	}

	return lexer.input[start_position:lexer.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
