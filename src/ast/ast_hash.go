package ast

import (
	"strings"

	"github.com/ape-lang/ape/src/token"
)

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }

func (hl *HashLiteral) String() string {
	var sb strings.Builder
	pairs := []string{}

	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	sb.WriteString("{")
	sb.WriteString(strings.Join(pairs, ", "))
	sb.WriteString("}")
	return sb.String()
}
