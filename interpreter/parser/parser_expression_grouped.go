package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()
	exp := p.parseExpression(LOWEST)

	if !p.advanceIfNext(token.PARENR) {
		return nil
	}
	return exp
}
