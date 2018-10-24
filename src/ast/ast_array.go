package ast

import (
	"ape/src/token"
	"strings"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

func (al *ArrayLiteral) String() string {
	var sb strings.Builder
	elements := []string{}

	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	sb.WriteString("[")
	sb.WriteString(strings.Join(elements, ", "))
	sb.WriteString("]")

	return sb.String()
}
