package ast

import (
	"strings"

	"github.com/ape-lang/ape/src/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var sb strings.Builder

	for _, s := range bs.Statements {
		sb.WriteString(s.String())
	}

	return sb.String()
}
