package operation

import "fmt"

// Operation contains the definition of an operation
type Operation struct {
	Name         string // human-readable name of the operation
	OperandSizes []int  // the size for each operand
}

// maps the Opcodes to the corresponding Operations
var operations = map[Opcode]*Operation{
	Pop: {"Pop", []int{}}, // Pop the value on top of the stack

	// Primitives
	Constant: {"Constant", []int{2}}, // Set a constant (to the given value)
	True:     {"True", []int{}},
	False:    {"False", []int{}},
	Null:     {"Null", []int{}},

	// Arithmetic
	Add: {"Add", []int{}},
	Sub: {"Sub", []int{}},
	Mul: {"Mul", []int{}},
	Div: {"Div", []int{}},

	// Comparison
	Equal:       {"Equal", []int{}},
	NotEqual:    {"NotEqual", []int{}},
	GreaterThan: {"GreaterThan", []int{}},

	// Prefix/Infix
	Minus: {"Minus", []int{}},
	Bang:  {"Bang", []int{}},

	// Jumps
	Jump:          {"Jump", []int{2}},          // Jump if value on top of stack truthy (to the given instruction)
	JumpNotTruthy: {"JumpNotTruthy", []int{2}}, // Jump if value on top of stack not truthy (to the given instruction)

	// Variables
	GetGlobal: {"GetGlobal", []int{2}}, // Get a Global variable definition (at the given index)
	SetGlobal: {"GetGlobal", []int{2}}, // Set a Global variable definition (with the given value)
	GetLocal:  {"GetLocal", []int{1}},  // Get a Local variable definition (at the given index)
	SetLocal:  {"GetLocal", []int{1}},  // Set a Local variable definition (with the given value)

	// Data Structures
	Array: {"Array", []int{2}}, // Create an Array literal (with the given declaration)
	Hash:  {"Hash", []int{2}},  // Create a Hash literal (with the given declaration)
	Index: {"Index", []int{}},  // Index operator

	// Functions
	Call:        {"Call", []int{1}},       // Call the function on top of the stack (with the given argument count)
	Return:      {"Return", []int{}},      // Return nothing, exit the function and return nil
	ReturnValue: {"ReturnValue", []int{}}, // Returns the value on top of the stack

	// Builtins
	GetBuiltin: {"GetBuiltin", []int{1}}, // Get a Builtin definition

	// Closures
	Closure: {"Closure", []int{2, 1}}, // The first operand references the function, the second the free variable count (max 1 byte)
	GetFree: {"GetFree", []int{1}},
}

// Lookup looks up a given Opcode and returns the corresponding Operation
func Lookup(value byte) (*Operation, error) {
	opcode := Opcode(value)
	operation, ok := operations[opcode]

	if !ok {
		return nil, fmt.Errorf("Undefined Opcode: %d", opcode)
	}
	return operation, nil
}
