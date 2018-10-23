package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalIdentifier(
	node *ast.Identifier,
	env *data.Environment,
) data.Data {
	val, ok := env.Get(node.Value)
	if !ok {
		return evalError("Identifier not found: " + node.Value)
	}
	return val
}
