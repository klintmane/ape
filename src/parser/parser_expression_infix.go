package parser

import "ape/src/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.current,
		Operator: p.current.Literal,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.advance()
	expression.Right = p.parseExpression(precedence)

	return expression
}
