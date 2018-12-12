package operation

// Opcode is the opcode value type, ex. '0000 0000'
type Opcode byte

// the available opcode values
const (
	Pop Opcode = iota

	Constant

	Add
	Sub
	Mul
	Div

	True
	False

	Equal
	NotEqual
	GreaterThan

	Minus
	Bang

	Jump
	JumpNotTruthy

	Null

	GetGlobal
	SetGlobal

	Array
	Hash

	Index
)
