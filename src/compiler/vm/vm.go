package vm

import (
	"ape/src/compiler/compiler"
	"ape/src/compiler/operation"
	"ape/src/data"
)

// VM contains the definition of the VM
type VM struct {
	constants    []data.Data
	instructions operation.Instruction
	stack        *Stack
}

// New creates a new VM from the given Bytecode
func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        NewStack(2048),
	}
}

// Result returns the value of the last popped element from the stack (last evaluated expression)
func (vm *VM) Result() data.Data {
	return vm.stack.popped()
}

// Run executes every instruction given to the VM on creation
func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := operation.Opcode(vm.instructions[ip])

		switch op {
		case operation.Pop:
			vm.stack.pop()

		case operation.Constant:
			constIndex := operation.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.stack.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case operation.Add:
			right := vm.stack.pop()
			left := vm.stack.pop()

			leftValue := left.(*data.Integer).Value
			rightValue := right.(*data.Integer).Value

			result := leftValue + rightValue
			vm.stack.push(&data.Integer{Value: result})
		}
	}
	return nil
}
