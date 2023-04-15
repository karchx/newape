package tokens

// TokenType type of token
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"

	PLUS = "PLUS"

	NUM  = "NUM" // decimal number
)

// Token is a token returned by the lexer
type Token struct {
	Type    TokenType
	Literal string
}
