package vm

func (vm *VM) executeBangOp() error {
	operand := vm.stack.pop()

	switch operand {
	case TRUE:
		return vm.stack.push(FALSE)
	case FALSE:
		return vm.stack.push(TRUE)
	default:
		return vm.stack.push(FALSE)
	}
}
