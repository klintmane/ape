package parser

import (
	"ape/src/ast"
	"ape/src/token"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParsers[p.current.Type]

	if prefix == nil {
		p.prefixParserError(p.current.Type)
		return nil
	}

	leftExp := prefix()

	for !p.isNext(token.SEMICOLON) && precedence < p.nextPrecedence() {
		infix := p.infixParsers[p.next.Type]

		if infix == nil {
			return leftExp
		}

		p.advance()
		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.current}
	stmt.Expression = p.parseExpression(LOWEST)

	if p.isNext(token.SEMICOLON) {
		p.advance()
	}
	return stmt
}

func (p *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if p.isNext(end) {
		p.advance()
		return list
	}

	p.advance()
	list = append(list, p.parseExpression(LOWEST))

	for p.isNext(token.COMMA) {
		p.advance()
		p.advance()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.advanceIfNext(end) {
		return nil
	}

	return list
}
