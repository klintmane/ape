package operation

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{SetConstant, []int{65534}, []byte{byte(SetConstant), 255, 254}},
	}

	for _, test := range tests {
		instruction := Instruction(test.opcode, test.operands...)

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
