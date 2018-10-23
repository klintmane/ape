package eval

import (
	"ape/interpreter/ast"
	"ape/interpreter/data"
)

func evalIfExpression(ie *ast.IfExpression) data.Data {
	condition := Eval(ie.Condition)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ie.Consequent)
	} else if ie.Alternate != nil {
		return Eval(ie.Alternate)
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
