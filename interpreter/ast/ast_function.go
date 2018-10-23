package ast

import (
	"ape/interpreter/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

func (fl *FunctionLiteral) String() string {
	var sb strings.Builder
	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	sb.WriteString(fl.TokenLiteral())
	sb.WriteString("(")
	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") ")
	sb.WriteString(fl.Body.String())

	return sb.String()
}
