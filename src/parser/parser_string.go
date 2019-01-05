package parser

import "github.com/ape-lang/ape/src/ast"

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.current, Value: p.current.Literal}
}
