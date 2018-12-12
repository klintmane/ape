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

func testBooleanData(expected bool, actual data.Data) error {
	result, ok := actual.(*data.Boolean)
	if !ok {
		return fmt.Errorf("data is not Boolean. got=%T (%+v)",
			actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("data has wrong value. got=%t, want=%t",
			result.Value, expected)
	}
	return nil
}

func testStringData(expected string, actual data.Data) error {
	result, ok := actual.(*data.String)
	if !ok {
		return fmt.Errorf("data is not String. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("data has wrong value. got=%q, want=%q", result.Value, expected)
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

func testExpectedData(t *testing.T, expected interface{}, actual data.Data) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerData(int64(expected), actual)

		if err != nil {
			t.Errorf("testIntegerData failed: %s", err)
		}

	case bool:
		err := testBooleanData(bool(expected), actual)
		if err != nil {
			t.Errorf("testBooleanData failed: %s", err)
		}

	case *data.Null:
		if actual != NULL {
			t.Errorf("data is not Null: %T (%+v)", actual, actual)
		}

	case string:
		err := testStringData(expected, actual)
		if err != nil {
			t.Errorf("testStringData failed: %s", err)
		}

	case []int:
		array, ok := actual.(*data.Array)
		if !ok {
			t.Errorf("data not Array: %T (%+v)", actual, actual)
			return
		}
		if len(array.Elements) != len(expected) {
			t.Errorf("wrong num of elements. want=%d, got=%d", len(expected), len(array.Elements))
			return
		}
		for i, expectedElem := range expected {
			err := testIntegerData(int64(expectedElem), array.Elements[i])
			if err != nil {
				t.Errorf("testIntegerData failed: %s", err)
			}
		}

	case map[data.HashKey]int64:
		hash, ok := actual.(*data.Hash)
		if !ok {
			t.Errorf("object is not Hash. got=%T (%+v)", actual, actual)
			return
		}

		if len(hash.Pairs) != len(expected) {
			t.Errorf("hash has wrong number of Pairs. want=%d, got=%d", len(expected), len(hash.Pairs))
			return
		}
		for expectedKey, expectedValue := range expected {
			pair, ok := hash.Pairs[expectedKey]
			if !ok {
				t.Errorf("no pair for given key in Pairs")
			}

			err := testIntegerData(expectedValue, pair.Value)
			if err != nil {
				t.Errorf("testIntegerData failed: %s", err)
			}
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
		{"-5", -5},
		{"-10", -10},
		{"-50 + 100 + -50", 0},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	runVMTests(t, tests)
}

func TestBooleanExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!(if (false) { 5; })", true},
	}

	runVMTests(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []vmTestCase{
		{"if (true) { 10 }", 10},
		{"if (true) { 10 } else { 20 }", 10},
		{"if (false) { 10 } else { 20 } ", 20},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 > 2) { 10 }", NULL},
		{"if (false) { 10 }", NULL},
		{"if ((if (false) { 10 })) { 10 } else { 20 }", 20},
	}

	runVMTests(t, tests)
}

func TestGlobalLetStatements(t *testing.T) {
	tests := []vmTestCase{
		{"let one = 1; one", 1},
		{"let one = 1; let two = 2; one + two", 3},
		{"let one = 1; let two = one + one; one + two", 3},
	}

	runVMTests(t, tests)
}

func TestStringExpressions(t *testing.T) {
	tests := []vmTestCase{
		{`"apelang"`, "apelang"},
		{`"ape" + "lang"`, "apelang"},
		{`"ape" + "lang" + "uage"`, "apelanguage"},
	}

	runVMTests(t, tests)
}

func TestArrayLiterals(t *testing.T) {
	tests := []vmTestCase{
		{"[]", []int{}},
		{"[1, 2, 3]", []int{1, 2, 3}},
		{"[1 + 2, 3 * 4, 5 + 6]", []int{3, 12, 11}},
	}
	runVMTests(t, tests)
}

func TestHashLiterals(t *testing.T) {
	tests := []vmTestCase{
		{
			"{}", map[data.HashKey]int64{},
		},
		{
			"{1: 2, 2: 3}",
			map[data.HashKey]int64{
				data.HashData(&data.Integer{Value: 1}): 2,
				data.HashData(&data.Integer{Value: 2}): 3,
			},
		},
		{
			"{1 + 1: 2 * 2, 3 + 3: 4 * 4}",
			map[data.HashKey]int64{
				data.HashData(&data.Integer{Value: 2}): 4,
				data.HashData(&data.Integer{Value: 6}): 16,
			},
		},
	}
	runVMTests(t, tests)
}
