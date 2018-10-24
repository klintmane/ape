package ast

import (
	"ape/src/token"
	"strings"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }

func (oe *InfixExpression) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(oe.Left.String())
	sb.WriteString(" " + oe.Operator + " ")
	sb.WriteString(oe.Right.String())
	sb.WriteString(")")

	return sb.String()
}
