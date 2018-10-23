package parser

import "ape/interpreter/token"

// Operators listed by lower precedence
const (
	_ int = iota // Declares the constants in the block as incrementing ints
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

// Precedence mapping to token
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NEQ:      EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
	token.PARENL:   CALL,
}

func precedence(t token.Token) int {
	if p, ok := precedences[t.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) nextPrecedence() int {
	return precedence(p.next)
}

func (p *Parser) currentPrecedence() int {
	return precedence(p.current)
}
