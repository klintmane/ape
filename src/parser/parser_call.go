package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.current, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.isNext(token.PARENR) {
		p.advance()
		return args
	}

	p.advance()
	args = append(args, p.parseExpression(LOWEST))

	for p.isNext(token.COMMA) {
		p.advance()
		p.advance()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.advanceIfNext(token.PARENR) {
		return nil
	}

	return args
}
