package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.current}
	array.Elements = p.parseExpressionList(token.BRACKETR)
	return array
}
