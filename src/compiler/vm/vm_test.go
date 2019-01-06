package vm

import (
	"fmt"
	"testing"

	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/compiler/compiler"
	"github.com/ape-lang/ape/src/data"
	"github.com/ape-lang/ape/src/lexer"
	"github.com/ape-lang/ape/src/parser"
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
		if actual != data.NULL {
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

	case *data.Error:
		err, ok := actual.(*data.Error)
		if !ok {
			t.Errorf("object is not Error: %T (%+v)", actual, actual)
			return
		}
		if err.Message != expected.Message {
			t.Errorf("wrong error message. expected=%q, got=%q", expected.Message, err.Message)
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
		{"if (1 > 2) { 10 }", data.NULL},
		{"if (false) { 10 }", data.NULL},
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

func TestIndexExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][0 + 2]", 3},
		{"[[1, 1, 1]][0][0]", 1},
		{"[][0]", data.NULL},
		{"[1, 2, 3][99]", data.NULL},
		{"[1][-1]", data.NULL},
		{"{1: 1, 2: 2}[1]", 1},
		{"{1: 1, 2: 2}[2]", 2},
		{"{1: 1}[0]", data.NULL},
		{"{}[0]", data.NULL},
	}
	runVMTests(t, tests)
}

func TestCallingFunctionsWithoutArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let eleven = fn() { 6 + 5; };
				eleven();
			`,
			expected: 11,
		},
		{
			input: `
				let one = fn() { 1; };
				let two = fn() { 2; };
				one() + two()
			`,
			expected: 3,
		},
		{
			input: `
				let a = fn() { 1 };
				let b = fn() { a() + 1 };
				let c = fn() { b() + 1 };
				c();
			`,
			expected: 3,
		},
	}
	runVMTests(t, tests)
}

func TestFunctionsWithReturnStatement(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let earlyReturn = fn() { return 99; 100; };
				earlyReturn();
			`,
			expected: 99,
		},
		{
			input: `
				let earlyReturn = fn() { return 99; return 100; };
				earlyReturn();
			`,
			expected: 99,
		},
	}
	runVMTests(t, tests)
}

func TestFunctionsWithoutReturnValue(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let noReturn = fn() { };
				noReturn();
			`,
			expected: data.NULL,
		},
		{
			input: `
				let noReturn = fn() { };
				let noReturnTwo = fn() { noReturn(); };
				noReturn();
				noReturnTwo();
			`,
			expected: data.NULL,
		},
	}
	runVMTests(t, tests)
}

func TestFirstClassFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let returnsOne = fn() { 1; };
				let returnsOneReturner = fn() { returnsOne; };
				returnsOneReturner()();
			`,
			expected: 1,
		},
		{
			input: `
			let returnsOneReturner = fn() {
				let returnsOne = fn() { 1; };
				returnsOne;
			};
			returnsOneReturner()();
			`,
			expected: 1,
		},
	}
	runVMTests(t, tests)
}

func TestCallingFunctionsWithBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let one = fn() { let one = 1; one };
				one();
			`,
			expected: 1,
		},
		{
			input: `
				let oneAndTwo = fn() { let one = 1; let two = 2; one + two; };
				oneAndTwo();
			`,
			expected: 3,
		},
		{
			input: `
				let oneAndTwo = fn() { let one = 1; let two = 2; one + two; };
				let threeAndFour = fn() { let three = 3; let four = 4; three + four; };
				oneAndTwo() + threeAndFour();
			`,
			expected: 10,
		},
		{
			input: `
				let firstFoobar = fn() { let foobar = 50; foobar; };
				let secondFoobar = fn() { let foobar = 100; foobar; };
				firstFoobar() + secondFoobar();
			`,
			expected: 150,
		},
		{
			input: `
				let globalSeed = 50;
				let minusOne = fn() {
					let num = 1;
					globalSeed - num;
				}
				let minusTwo = fn() {
					let num = 2;
					globalSeed - num;
				}
				minusOne() + minusTwo();
			`,
			expected: 97,
		},
	}
	runVMTests(t, tests)
}

func TestCallingFunctionsWithArgumentsAndBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let identity = fn(a) { a; };
				identity(4);
			`,
			expected: 4,
		},
		{
			input: `
				let sum = fn(a, b) { a + b; };
				sum(1, 2);
			`,
			expected: 3,
		},

		{
			input: `
			let sum = fn(a, b) {
			let c = a + b;
				c;
			};
			sum(1, 2);
			`,
			expected: 3,
		},
		{
			input: `
			let sum = fn(a, b) {
				let c = a + b;
				c;
			};
			sum(1, 2) + sum(3, 4);`,
			expected: 10,
		},
		{
			input: `
				let sum = fn(a, b) {
					let c = a + b;
					c;
				};
				let outer = fn() {
					sum(1, 2) + sum(3, 4);
				};
				outer();
			`,
			expected: 10,
		},
		{
			input: `
			let globalNum = 10;
			let sum = fn(a, b) {
				let c = a + b;
				c + globalNum;
			};
			let outer = fn() {
				sum(1, 2) + sum(3, 4) + globalNum;
			};
			outer() + globalNum;
			`,
			expected: 50,
		},
	}
	runVMTests(t, tests)
}

func TestCallingFunctionsWithWrongArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input:    `fn() { 1; }(1);`,
			expected: `wrong number of arguments: want=0, got=1`,
		},
		{
			input:    `fn(a) { a; }();`,
			expected: `wrong number of arguments: want=1, got=0`,
		},
		{
			input:    `fn(a, b) { a + b; }(1);`,
			expected: `wrong number of arguments: want=2, got=1`,
		},
	}

	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err == nil {
			t.Fatalf("expected VM error but resulted in none.")
		}

		if err.Error() != tt.expected {
			t.Fatalf("wrong VM error: want=%q, got=%q", tt.expected, err)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []vmTestCase{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{
			`len(1)`,
			&data.Error{Message: "argument to 'len' not supported, got INTEGER"},
		},
		{`len("one", "two")`,
			&data.Error{Message: "wrong number of arguments. got=2, want=1"},
		},
		{`len([1, 2, 3])`, 3},
		{`len([])`, 0},
		{`print("hello", "world!")`, data.NULL},
		{`head([1, 2, 3])`, 1},
		{`head([])`, data.NULL},
		{`head(1)`,
			&data.Error{Message: "argument to 'first' must be ARRAY, got INTEGER"},
		},
		{`last([1, 2, 3])`, 3},
		{`last([])`, data.NULL},
		{`last(1)`,
			&data.Error{Message: "argument to 'last' must be ARRAY, got INTEGER"},
		},
		{`tail([1, 2, 3])`, []int{2, 3}},
		{`tail([])`, data.NULL},
		{`push([], 1)`, []int{1}},
		{`push(1, 1)`,
			&data.Error{Message: "argument to 'push' must be ARRAY, got INTEGER"},
		},
	}
	runVMTests(t, tests)
}
