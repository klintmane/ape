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
	}
	runCompilerTests(t, tests)
}

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

// * HELPERS

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
		}
	}
	return nil
}

func testIntegerData(expected int64, actual data.Data) error {
	result, ok := actual.(*data.Integer)

	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. got=%d, want=%d", result.Value, expected)
	}
	return nil
}
