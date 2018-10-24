package eval

import (
	"ape/src/ast"
	"ape/src/interpreter/data"
)

func evalIdentifier(node *ast.Identifier, env *data.Environment) data.Data {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := evalBuiltin(node.Value); ok {
		return builtin
	}

	return evalError("Identifier not found: " + node.Value)
}
