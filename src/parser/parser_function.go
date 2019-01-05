package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.current}
	if !p.advanceIfNext(token.PARENL) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()
	if !p.advanceIfNext(token.BRACEL) {
		return nil
	}

	lit.Body = p.parseBlockStatement()
	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if p.isNext(token.PARENR) {
		p.advance()
		return identifiers
	}

	p.advance()
	ident := &ast.Identifier{Token: p.current, Value: p.current.Literal}
	identifiers = append(identifiers, ident)

	for p.isNext(token.COMMA) {
		p.advance()
		p.advance()
		ident := &ast.Identifier{Token: p.current, Value: p.current.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.advanceIfNext(token.PARENR) {
		return nil
	}

	return identifiers
}
