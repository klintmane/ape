package ast

import (
	"strings"

	"github.com/ape-lang/ape/src/token"
)

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *IndexExpression) String() string {
	var sb strings.Builder

	sb.WriteString("(")
	sb.WriteString(ie.Left.String())
	sb.WriteString("[")
	sb.WriteString(ie.Index.String())
	sb.WriteString("])")

	return sb.String()
}
