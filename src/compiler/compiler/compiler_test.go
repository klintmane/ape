package compiler

import (
	"ape/src/ast"
	"ape/src/compiler/operation"
	"ape/src/data"
	"ape/src/lexer"
	"ape/src/parser"
	"fmt"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []operation.Instruction
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Add),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1; 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Pop),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 - 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Sub),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 * 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Mul),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "2 / 1",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Div),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "-1",
			expectedConstants: []interface{}{1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Minus),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestBooleanExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "true",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.True),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "false",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.False),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 > 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.GreaterThan),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 < 2",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.GreaterThan),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 == 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Equal),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "1 != 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.NotEqual),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "true == false",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.True),
				operation.NewInstruction(operation.False),
				operation.NewInstruction(operation.Equal),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "true != false",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.True),
				operation.NewInstruction(operation.False),
				operation.NewInstruction(operation.NotEqual),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "!true",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.True),
				operation.NewInstruction(operation.Bang),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             `if (true) { 10 }; 3333;`,
			expectedConstants: []interface{}{10, 3333},
			expectedInstructions: []operation.Instruction{
				// 0000
				operation.NewInstruction(operation.True),
				// 0001
				operation.NewInstruction(operation.JumpNotTruthy, 10),
				// 0004
				operation.NewInstruction(operation.Constant, 0),
				// 0007
				operation.NewInstruction(operation.Jump, 11),
				// 0010
				operation.NewInstruction(operation.Null),
				// 0011
				operation.NewInstruction(operation.Pop),
				// 0012
				operation.NewInstruction(operation.Constant, 1),
				// 0015
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             `if (true) { 10 } else { 20 }; 3333;`,
			expectedConstants: []interface{}{10, 20, 3333},
			expectedInstructions: []operation.Instruction{
				// 0000
				operation.NewInstruction(operation.True),
				// 0001
				operation.NewInstruction(operation.JumpNotTruthy, 10),
				// 0004
				operation.NewInstruction(operation.Constant, 0),
				// 0007
				operation.NewInstruction(operation.Jump, 13),
				// 0010
				operation.NewInstruction(operation.Constant, 1),
				// 0013
				operation.NewInstruction(operation.Pop),
				// 0014
				operation.NewInstruction(operation.Constant, 2),
				// 0017
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestGlobalLetStatements(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
				let one = 1;
				let two = 2;
			`,
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.SetGlobal, 1),
			},
		},
		{
			input: `
				let one = 1;
				one;
			`,
			expectedConstants: []interface{}{1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				let one = 1;
				let two = one;
				two;
			`,
			expectedConstants: []interface{}{1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.SetGlobal, 1),
				operation.NewInstruction(operation.GetGlobal, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestStringExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             `"apelang"`,
			expectedConstants: []interface{}{"apelang"},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             `"ape" + "lang"`,
			expectedConstants: []interface{}{"ape", "lang"},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Add),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestArrayLiterals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "[]",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Array, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "[1, 2, 3]",
			expectedConstants: []interface{}{1, 2, 3},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Array, 3),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "[1 + 2, 3 - 4, 5 * 6]",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Add),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Sub),
				operation.NewInstruction(operation.Constant, 4),
				operation.NewInstruction(operation.Constant, 5),
				operation.NewInstruction(operation.Mul),
				operation.NewInstruction(operation.Array, 3),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

// * HELPERS

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
		compiler := New()
		err := compiler.Compile(program)

		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(flattenInstructions(tt.expectedInstructions), bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %s", err)
		}
	}
}

// lexes and parses a program, returning an AST
func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)

	return p.ParseProgram()
}

// flattens the test expected instructions slice so it's comparable to the bytecode instructions slice
func flattenInstructions(instructions []operation.Instruction) operation.Instruction {
	result := operation.Instruction{}

	for _, current := range instructions {
		result = append(result, current...)
	}
	return result
}

// tests expected instructions
func testInstructions(expected operation.Instruction, actual operation.Instruction) error {
	if len(actual) != len(expected) {
		return fmt.Errorf("wrong instructions length.\nwant=%q\ngot =%q", expected, actual)
	}

	for i, ins := range expected {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot =%q", i, expected, actual)
		}
	}
	return nil
}

// tests expected constants
func testConstants(t *testing.T, expected []interface{}, actual []data.Data) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. got=%d, want=%d",
			len(actual), len(expected))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerData(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testIntegerData failed: %s", i, err)
			}

		case string:
			err := testStringData(constant, actual[i])
			if err != nil {
				return fmt.Errorf("constant %d - testStringData failed: %s", i, err)
			}
		}
	}
	return nil
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
