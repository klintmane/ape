package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
	"fmt"
)

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.current,
		Operator: p.current.Literal,
	}

	p.advance()
	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) prefixParserError(t token.TokenType) {
	msg := fmt.Sprintf("No prefix parser found for %s", t)
	p.errors = append(p.errors, msg)
}
