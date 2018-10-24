package parser

import (
	"ape/src/ast"
	"ape/src/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.current, Value: p.isCurrent(token.TRUE)}
}
