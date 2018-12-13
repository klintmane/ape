package operation

// Opcode is the opcode value type, ex. '0000 0000'
type Opcode byte

// the available opcode values
const (
	Pop Opcode = iota

	// Primitives
	Constant
	True
	False
	Null

	// Arithmetic
	Add
	Sub
	Mul
	Div

	// Comparison
	Equal
	NotEqual
	GreaterThan

	// Prefix/Infix
	Minus
	Bang

	// Jumps
	Jump
	JumpNotTruthy

	// Variables
	GetGlobal
	SetGlobal

	// Data Structures
	Array
	Hash
	Index

	// Functions
	Call
	Return
	ReturnValue
)
