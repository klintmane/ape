package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.current}
	if !p.advanceIfNext(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.current, Value: p.current.Literal}
	if !p.advanceIfNext(token.ASSIGN) {
		return nil
	}

	p.advance()
	stmt.Value = p.parseExpression(LOWEST)
	if p.isNext(token.SEMICOLON) {
		p.advance()
	}

	return stmt
}
