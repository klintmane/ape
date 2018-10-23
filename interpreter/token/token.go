package token

const (
	IDENT = "IDENT"

	FUNCTION = "FUNCTION"
	LET      = "LET"

	INT   = "INT"
	TRUE  = "TRUE"
	FALSE = "FALSE"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	BANG = "!"
	LT   = "<"
	GT   = ">"
	EQ   = "=="
	NEQ  = "!="

	PARENL    = "("
	PARENR    = ")"
	BRACEL    = "{"
	BRACER    = "}"
	COMMA     = ","
	SEMICOLON = ";"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(tokenType TokenType, char byte) Token {
	return Token{Type: tokenType, Literal: string(char)}
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,

	"true":  TRUE,
	"false": FALSE,

	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tType, ok := keywords[ident]; ok {
		return tType
	}
	return IDENT
}
