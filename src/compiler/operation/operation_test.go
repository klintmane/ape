package operation

import "testing"

// Tests that Operations are correctly converted into Instruction bytes
func TestNewInstruction(t *testing.T) {
	tests := []struct {
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{Constant, []int{65534}, []byte{byte(Constant), 255, 254}},
		{Add, []int{}, []byte{byte(Add)}},
		{GetLocal, []int{255}, []byte{byte(GetLocal), 255}},
		{Closure, []int{65534, 255}, []byte{byte(Closure), 255, 254, 255}},
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
		NewInstruction(Add),
		NewInstruction(GetLocal, 1),
		NewInstruction(Constant, 2),
		NewInstruction(Constant, 65535),
		NewInstruction(Closure, 65535, 255),
	}
	expected := `0000 Add
0001 GetLocal 1
0003 Constant 2
0006 Constant 65535
0009 Closure 65535 255
`

	result := Instruction{}

	for _, ins := range instructions {
		result = append(result, ins...)
	}

	if result.String() != expected {
		t.Errorf("instructions wrongly formatted.\nwant=%q\ngot=%q",
			expected, result.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{Constant, []int{65535}, 2},
		{GetLocal, []int{255}, 1},
		{Closure, []int{65535, 255}, 3},
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
