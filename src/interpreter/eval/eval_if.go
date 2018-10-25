package eval

import (
	"ape/src/ast"
	"ape/src/data"
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
		return NULL
	}
}

func isTruthy(d data.Data) bool {
	switch d {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
