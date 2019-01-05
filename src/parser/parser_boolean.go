package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.current, Value: p.isCurrent(token.TRUE)}
}
