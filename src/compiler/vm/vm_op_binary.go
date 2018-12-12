package vm

import (
	"ape/src/compiler/operation"
	"ape/src/data"
	"fmt"
)

func (vm *VM) executeBinaryOp(op operation.Opcode) error {
	right := vm.stack.pop()
	left := vm.stack.pop()

	leftType := left.Type()
	rightType := right.Type()

	switch {
	case leftType == data.INTEGER_TYPE && rightType == data.INTEGER_TYPE:
		return vm.executeBinaryIntegerOp(op, left, right)

	case leftType == data.STRING_TYPE && rightType == data.STRING_TYPE:
		return vm.executeBinaryStringOp(op, left, right)

	default:
		return fmt.Errorf("unsupported types for binary operation: %s %s", leftType, rightType)
	}
}

func (vm *VM) executeBinaryIntegerOp(op operation.Opcode, left, right data.Data) error {
	var result int64

	leftVal := left.(*data.Integer).Value
	rightVal := right.(*data.Integer).Value

	switch op {
	case operation.Add:
		result = leftVal + rightVal

	case operation.Sub:
		result = leftVal - rightVal

	case operation.Mul:
		result = leftVal * rightVal

	case operation.Div:
		result = leftVal / rightVal

	default:
		return fmt.Errorf("unknown integer operator: %d", op)
	}
	return vm.stack.push(&data.Integer{Value: result})
}

func (vm *VM) executeBinaryStringOp(op operation.Opcode, left, right data.Data) error {
	var result string

	leftVal := left.(*data.String).Value
	rightVal := right.(*data.String).Value

	switch op {
	case operation.Add:
		result = leftVal + rightVal

	default:
		return fmt.Errorf("unknown string operator: %d", op)
	}
	return vm.stack.push(&data.String{Value: result})
}
