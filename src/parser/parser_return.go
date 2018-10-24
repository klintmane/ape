package parser

import (
	"ape/src/ast"
	"ape/src/token"
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
