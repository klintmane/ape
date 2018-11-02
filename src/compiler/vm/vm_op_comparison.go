package vm

import (
	"ape/src/compiler/operation"
	"ape/src/data"
	"fmt"
)

func (vm *VM) executeComparison(op operation.Opcode) error {
	right := vm.stack.pop()
	left := vm.stack.pop()

	if left.Type() == data.INTEGER_TYPE || right.Type() == data.INTEGER_TYPE {
		return vm.executeIntegerComparison(op, left, right)
	}

	switch op {
	case operation.Equal:
		return vm.executeBoolean(right == left)

	case operation.NotEqual:
		return vm.executeBoolean(right != left)

	default:
		return fmt.Errorf("unknown operator: %d (%s %s)",
			op, left.Type(), right.Type())
	}
}

func (vm *VM) executeIntegerComparison(op operation.Opcode, left, right data.Data) error {
	leftValue := left.(*data.Integer).Value
	rightValue := right.(*data.Integer).Value

	switch op {
	case operation.Equal:
		return vm.executeBoolean(rightValue == leftValue)

	case operation.NotEqual:
		return vm.executeBoolean(rightValue != leftValue)

	case operation.GreaterThan:
		return vm.executeBoolean(leftValue > rightValue)

	default:
		return fmt.Errorf("unknown operator: %d", op)
	}
}

func (vm *VM) executeBoolean(val bool) error {
	if val {
		return vm.stack.push(TRUE)
	}
	return vm.stack.push(FALSE)
}
