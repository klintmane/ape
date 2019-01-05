package eval

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/data"
)

func evalFunction(node *ast.FunctionLiteral, env *data.Environment) *data.Function {
	params := node.Parameters
	body := node.Body

	return &data.Function{Parameters: params, Env: env, Body: body}
}
