package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.current}
	if !p.advanceIfNext(token.PARENL) {
		return nil
	}
	p.advance()

	expression.Condition = p.parseExpression(LOWEST)
	if !p.advanceIfNext(token.PARENR) {
		return nil
	}

	if !p.advanceIfNext(token.BRACEL) {
		return nil
	}

	expression.Consequent = p.parseBlockStatement()

	if p.isNext(token.ELSE) {
		p.advance()
		if !p.advanceIfNext(token.BRACEL) {
			return nil
		}
		expression.Alternate = p.parseBlockStatement()
	}

	return expression
}
