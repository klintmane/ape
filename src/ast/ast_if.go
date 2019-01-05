package ast

import (
	"strings"

	"github.com/ape-lang/ape/src/token"
)

type IfExpression struct {
	Token      token.Token
	Condition  Expression
	Consequent *BlockStatement
	Alternate  *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IfExpression) String() string {
	var sb strings.Builder

	sb.WriteString("if")
	sb.WriteString(ie.Condition.String())
	sb.WriteString(" ")
	sb.WriteString(ie.Consequent.String())

	if ie.Alternate != nil {
		sb.WriteString(" else ")
		sb.WriteString(ie.Alternate.String())
	}

	return sb.String()
}
