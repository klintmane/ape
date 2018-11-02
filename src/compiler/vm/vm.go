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
	for pointer := 0; pointer < len(vm.instructions); pointer++ {
		op := operation.Opcode(vm.instructions[pointer])

		switch op {
		case operation.Pop:
			vm.stack.pop()

		case operation.Constant:
			constIndex := operation.ReadUint16(vm.instructions[pointer+1:])
			pointer += 2

			err := vm.stack.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case operation.Add, operation.Sub, operation.Mul, operation.Div:
			err := vm.executeBinaryOp(op)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
