package operation

import "fmt"

// Lookup looks up a given opcode and returns the corresponding operation
func Lookup(value byte) (*Operation, error) {
	opcode := Opcode(value)
	def, ok := operations[opcode]

	if !ok {
		return nil, fmt.Errorf("Undefined Opcode: %d", opcode)
	}
	return def, nil
}
