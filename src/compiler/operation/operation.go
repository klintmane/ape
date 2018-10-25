package operation

// Operation contains the definition of an operation
type Operation struct {
	Name         string // human-readable name of the operation
	OperandSizes []int  // the size for each operand
}

// maps the opcodes to the corresponding operations
var operations = map[Opcode]*Operation{
	SetConstant: {"SetConstant", []int{2}},
}
