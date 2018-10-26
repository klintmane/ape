package operation

import "testing"

// Tests that Operations are correctly converted into Instruction bytes
func TestInstruction(t *testing.T) {
	tests := []struct {
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
	}

	for _, test := range tests {
		instruction := NewInstruction(test.opcode, test.operands...)

		if len(instruction) != len(test.expected) {
			t.Errorf("instruction has wrong length. want=%d, got=%d", len(test.expected), len(instruction))
		}

		for i, b := range test.expected {
			if instruction[i] != test.expected[i] {
				t.Errorf("wrong byte at pos %d. want=%d, got=%d", i, b, instruction[i])
			}
		}
	}
}

// Tests that Instruction correctly implement a String method and are correctly printed
func TestInstructionString(t *testing.T) {
	instructions := []Instruction{
		NewInstruction(OpConstant, 1),
		NewInstruction(OpConstant, 2),
		NewInstruction(OpConstant, 65535),
	}
	expected := `0000 OpConstant 1
0003 OpConstant 2
0006 OpConstant 65535
`

	concatted := Instruction{}

	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}

	if concatted.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q",
			expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
	}

	for _, tt := range tests {
		instruction := NewInstruction(tt.op, tt.operands...)
		def, err := Lookup(byte(tt.op))

		if err != nil {
			t.Fatalf("definition not found: %q\n", err)
		}

		operandsRead, n := ReadOperands(def, instruction[1:])

		if n != tt.bytesRead {
			t.Fatalf("n wrong. want=%d, got=%d", tt.bytesRead, n)
		}

		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("operand wrong. want=%d, got=%d", want, operandsRead[i])
			}
		}
	}
}
