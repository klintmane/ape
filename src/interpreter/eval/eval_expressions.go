package eval

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/data"
)

func evalExpressions(exps []ast.Expression, env *data.Environment) []data.Data {
	var result []data.Data

	for _, e := range exps {
		evaluated := Eval(e, env)

		if isError(evaluated) {
			return []data.Data{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}
