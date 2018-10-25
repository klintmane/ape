package operation

// Opcode is the opcode value type, ex. '0000 0000'
type Opcode byte

// the available opcode values
const (
	SetConstant Opcode = iota
)
