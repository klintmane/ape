package eval

import (
	"ape/src/ast"
	"ape/src/data"
)

func evalFunction(node *ast.FunctionLiteral, env *data.Environment) *data.Function {
	params := node.Parameters
	body := node.Body

	return &data.Function{Parameters: params, Env: env, Body: body}
}
