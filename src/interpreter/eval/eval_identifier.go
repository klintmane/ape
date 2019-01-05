package eval

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/data"
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
