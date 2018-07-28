package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	SLASH       = "/"
	ASTERISK    = "*"
	LESSTHAN    = "<"
	GREATERTHAN = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	NEWLINE   = "\n"

	// Keywords
	METHOD = "METHOD"
	END    = "END"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"def":    METHOD,
	"end":    END,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}
