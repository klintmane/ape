package vm

import (
	"ape/src/compiler/compiler"
	"ape/src/compiler/operation"
	"ape/src/data"
)

const GLOBALS_SIZE = 65536 // size of an operand

// Global references, so a new object does not get allocated for each evaluation
var (
	TRUE  = &data.Boolean{Value: true}
	FALSE = &data.Boolean{Value: false}
	NULL  = &data.Null{}
)

// VM contains the definition of the VM
type VM struct {
	instructions operation.Instruction
	constants    []data.Data
	globals      []data.Data
	stack        *Stack
}

// New creates a new VM from the given Bytecode
func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		globals:      make([]data.Data, GLOBALS_SIZE),
		stack:        NewStack(2048),
	}
}

func NewWithGlobals(bytecode *compiler.Bytecode, globals []data.Data) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		globals:      globals,
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

		case operation.True:
			err := vm.stack.push(TRUE)
			if err != nil {
				return err
			}

		case operation.False:
			err := vm.stack.push(FALSE)
			if err != nil {
				return err
			}

		case operation.Bang:
			err := vm.executeBangOp()
			if err != nil {
				return err
			}

		case operation.Minus:
			err := vm.executeMinusOp()
			if err != nil {
				return err
			}

		case operation.Add, operation.Sub, operation.Mul, operation.Div:
			err := vm.executeBinaryOp(op)
			if err != nil {
				return err
			}

		case operation.Equal, operation.NotEqual, operation.GreaterThan:
			err := vm.executeComparison(op)
			if err != nil {
				return err
			}

		case operation.Jump:
			pos := int(operation.ReadUint16(vm.instructions[pointer+1:]))
			pointer = pos - 1

		case operation.JumpNotTruthy:
			pos := int(operation.ReadUint16(vm.instructions[pointer+1:]))
			pointer += 2
			condition := vm.stack.pop()
			if !isTruthy(condition) {
				pointer = pos - 1
			}

		case operation.Null:
			err := vm.stack.push(NULL)
			if err != nil {
				return err
			}

		case operation.SetGlobal:
			index := operation.ReadUint16(vm.instructions[pointer+1:])
			pointer += 2
			vm.globals[index] = vm.stack.pop()

		case operation.GetGlobal:
			index := operation.ReadUint16(vm.instructions[pointer+1:])
			pointer += 2
			err := vm.stack.push(vm.globals[index])
			if err != nil {
				return err
			}

		case operation.Array:
			numElements := int(operation.ReadUint16(vm.instructions[pointer+1:]))
			pointer += 2
			array := vm.buildArray(vm.stack.pointer-numElements, vm.stack.pointer)
			vm.stack.pointer = vm.stack.pointer - numElements
			err := vm.stack.push(array)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isTruthy(d data.Data) bool {
	switch d := d.(type) {
	case *data.Boolean:
		return d.Value
	case *data.Null:
		return false
	default:
		return true
	}
}
