package vm

import (
	"ape/src/ast"
	"ape/src/compiler/compiler"
	"ape/src/data"
	"ape/src/lexer"
	"ape/src/parser"
	"fmt"
	"testing"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)

	return p.ParseProgram()
}

func testIntegerData(expected int64, actual data.Data) error {
	result, ok := actual.(*data.Integer)

	if !ok {
		return fmt.Errorf("data is not Integer. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("data has wrong value. got=%d, want=%d", result.Value, expected)
	}

	return nil
}

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVMTests(t *testing.T, tests []vmTestCase) {
	t.Helper()
	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)

		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()

		if err != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.stack.popped()
		testExpectedData(t, tt.expected, stackElem)
	}

}

func testExpectedData(
	t *testing.T,
	expected interface{},
	actual data.Data,
) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerData(int64(expected), actual)

		if err != nil {
			t.Errorf("testIntegerData failed: %s", err)
		}
	}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"50 / 2 * 2 + 10 - 5", 55},
		{"5 * (2 + 10)", 60},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * (2 + 10)", 60},
	}

	runVMTests(t, tests)
}
