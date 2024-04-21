package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL signifies a token/character we don’t know about
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTIOn"
	LET      = "LET"
)
