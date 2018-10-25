package operation

import "encoding/binary"

// Instruction takes an opcode and some operands and returns an instruction (array of bytes)
func Instruction(opcode Opcode, operands ...int) []byte {

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
	// Set offset to point to the next element
	instruction := make([]byte, instructionSize)
	instruction[0] = byte(opcode)
	offset := 1

	// For each operand
	for i, operand := range operands {
		// Get the operand size
		size := operation.OperandSizes[i]

		switch size {
		case 2:
			// Cast the operand to uint16 and put it onto the instruction array (in big-endian encoding)
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}

		// Set offset to point to the next element
		offset += size
	}
	return instruction
}
