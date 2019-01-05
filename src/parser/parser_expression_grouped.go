package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.advance()
	exp := p.parseExpression(LOWEST)

	if !p.advanceIfNext(token.PARENR) {
		return nil
	}
	return exp
}
