package parser

import (
	"ape/src/ast"
	"ape/src/token"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()
	exp := p.parseExpression(LOWEST)

	if !p.advanceIfNext(token.PARENR) {
		return nil
	}
	return exp
}
