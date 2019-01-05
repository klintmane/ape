package parser

import (
	"fmt"
	"strconv"

	"github.com/ape-lang/ape/src/ast"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.current}
	value, err := strconv.ParseInt(p.current.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("'%q' could not be parsed into an integer", p.current.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}
