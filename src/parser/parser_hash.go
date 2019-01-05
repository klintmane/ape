package parser

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/token"
)

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.current}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.isNext(token.BRACER) {
		p.advance()
		key := p.parseExpression(LOWEST)

		if !p.advanceIfNext(token.COLON) {
			return nil
		}

		p.advance()
		value := p.parseExpression(LOWEST)
		hash.Pairs[key] = value

		if !p.isNext(token.BRACER) && !p.advanceIfNext(token.COMMA) {
			return nil
		}
	}

	if !p.advanceIfNext(token.BRACER) {
		return nil
	}

	return hash
}
