package operation

import (
	"bytes"
	"fmt"
)

// String returns a human-readable version of the instruction
func (ins Instruction) String() string {
	var out bytes.Buffer
	i := 0

	for i < len(ins) {
		operation, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(operation, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(operation, operands))
		i += 1 + read
	}
	return out.String()
}

func (ins Instruction) fmtInstruction(operation *Operation, operands []int) string {
	operandCount := len(operation.OperandSizes)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return operation.Name
	case 1:
		return fmt.Sprintf("%s %d", operation.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", operation.Name, operands[0], operands[1])
	}
	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", operation.Name)
}
