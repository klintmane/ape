package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalFunction(node *ast.FunctionLiteral, env *data.Environment) *data.Function {
	params := node.Parameters
	body := node.Body

	return &data.Function{Parameters: params, Env: env, Body: body}
}
