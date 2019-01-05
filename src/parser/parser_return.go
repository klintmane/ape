package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
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
