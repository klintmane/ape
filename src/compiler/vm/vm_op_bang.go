package vm

import "github.com/ape-lang/ape/src/data"

func (vm *VM) executeBangOp() error {
	operand := vm.stack.pop()

	switch operand {
	case data.TRUE:
		return vm.stack.push(data.FALSE)
	case data.FALSE:
		return vm.stack.push(data.TRUE)
	case data.NULL:
		return vm.stack.push(data.TRUE)
	default:
		return vm.stack.push(data.FALSE)
	}
}
