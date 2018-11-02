package vm

import (
	"ape/src/data"
	"fmt"
)

func (vm *VM) executeMinusOp() error {
	operand := vm.stack.pop()

	if operand.Type() != data.INTEGER_TYPE {
		return fmt.Errorf("unsupported type for negation: %s", operand.Type())
	}

	value := operand.(*data.Integer).Value
	return vm.stack.push(&data.Integer{Value: -value})
}
