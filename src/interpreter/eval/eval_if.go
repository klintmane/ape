package eval

import (
	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/data"
)

func evalIfExpression(ie *ast.IfExpression, env *data.Environment) data.Data {
	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequent, env)
	} else if ie.Alternate != nil {
		return Eval(ie.Alternate, env)
	} else {
		return data.NULL
	}
}

func isTruthy(d data.Data) bool {
	switch d {
	case data.NULL:
		return false
	case data.TRUE:
		return true
	case data.FALSE:
		return false
	default:
		return true
	}
}
