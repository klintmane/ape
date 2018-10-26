package operation

import "fmt"

// Operation contains the definition of an operation
type Operation struct {
	Name         string // human-readable name of the operation
	OperandSizes []int  // the size for each operand
}

// maps the Opcodes to the corresponding Operations
var operations = map[Opcode]*Operation{
	OpConstant: {"OpConstant", []int{2}},
}

// Lookup looks up a given Opcode and returns the corresponding Operation
func Lookup(value byte) (*Operation, error) {
	opcode := Opcode(value)
	def, ok := operations[opcode]

	if !ok {
		return nil, fmt.Errorf("Undefined Opcode: %d", opcode)
	}
	return def, nil
}
