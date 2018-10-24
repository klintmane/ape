package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
)

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: p.current, Left: left}
	p.advance()
	exp.Index = p.parseExpression(LOWEST)

	if !p.advanceIfNext(token.BRACKETR) {
		return nil
	}

	return exp
}
