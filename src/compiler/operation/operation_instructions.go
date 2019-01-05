package operation

import (
	"encoding/binary"
)

// Instruction is the byte slice resulting from an operation
type Instruction []byte

// NewInstruction takes an opcode and some operands and returns an instruction (array of bytes)
func NewInstruction(opcode Opcode, operands ...int) Instruction {

	// Get the operation that corresponds to the given opcode
	operation, ok := operations[opcode]
	if !ok {
		return []byte{}
	}

	// Calculate the instruction size based on the size of the operation operands
	instructionSize := 1
	for _, size := range operation.OperandSizes {
		instructionSize += size
	}

	// Create a byte array with the instructionSize to hold the instruction
	// Set the opcode as the first element
	// Set position to point to the next element
	instruction := make([]byte, instructionSize)
	instruction[0] = byte(opcode)

	position := 1

	// For each operand
	for i, operand := range operands {
		// Get the operand size
		size := operation.OperandSizes[i]

		switch size {
		case 1:
			// No magic required for 1-byte sized operands
			instruction[position] = byte(operand)
		case 2:
			// Cast the operand to uint16 and put it onto the instruction array (in big-endian encoding)
			binary.BigEndian.PutUint16(instruction[position:], uint16(operand))
		}

		// Set position to point to the next element
		position += size
	}
	return instruction

}

// ReadOperands take an operation and instructions and returns the operands and the position
func ReadOperands(operation *Operation, ins Instruction) ([]int, int) {
	operands := make([]int, len(operation.OperandSizes))
	position := 0

	// For each operand size of the operation
	// Get the instruction slice corresponding to it and convert it to an int
	// Set the position to point to the next operand
	for i, size := range operation.OperandSizes {
		switch size {
		case 1:
			operands[i] = int(ReadUint8(ins[position:]))
		case 2:
			operands[i] = int(ReadUint16(ins[position:]))
		}
		position += size
	}
	return operands, position
}

// ReadUint16 returns the uint16 representation of an instruction
func ReadUint16(ins Instruction) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// ReadUint8 returns the uint8 representation of an instruction
func ReadUint8(ins Instruction) uint8 { return uint8(ins[0]) }
