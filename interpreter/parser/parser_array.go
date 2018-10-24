package parser

import (
	"ape/interpreter/ast"
	"ape/interpreter/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.current}
	array.Elements = p.parseExpressionList(token.BRACKETR)
	return array
}
