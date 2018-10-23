package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
)

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.current}
	p.advance()
	stmt.ReturnValue = p.parseExpression(LOWEST)

	if p.isNext(token.SEMICOLON) {
		p.advance()
	}

	return stmt
}
