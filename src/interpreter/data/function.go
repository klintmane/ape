package data

import (
	"ape/src/ast"
	"strings"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() DataType { return FUNCTION_TYPE }

func (f *Function) Inspect() string {
	var sb strings.Builder
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	sb.WriteString("fn")
	sb.WriteString("(")
	sb.WriteString(strings.Join(params, ", "))
	sb.WriteString(") {\n")
	sb.WriteString(f.Body.String())
	sb.WriteString("\n}")

	return sb.String()
}
