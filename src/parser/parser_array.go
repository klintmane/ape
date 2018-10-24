package parser

import (
	"ape/src/ast"
	"ape/src/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.current}
	array.Elements = p.parseExpressionList(token.BRACKETR)
	return array
}
