package compiler

import (
	"fmt"
	"testing"

	"github.com/ape-lang/ape/src/ast"
	"github.com/ape-lang/ape/src/compiler/operation"
	"github.com/ape-lang/ape/src/data"
	"github.com/ape-lang/ape/src/lexer"
	"github.com/ape-lang/ape/src/parser"
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

func TestHashLiterals(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "{}",
			expectedConstants: []interface{}{},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Hash, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "{1: 2, 3: 4, 5: 6}",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Constant, 4),
				operation.NewInstruction(operation.Constant, 5),
				operation.NewInstruction(operation.Hash, 6),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "{1: 2 + 3, 4: 5 * 6}",
			expectedConstants: []interface{}{1, 2, 3, 4, 5, 6},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Add),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Constant, 4),
				operation.NewInstruction(operation.Constant, 5),
				operation.NewInstruction(operation.Mul),
				operation.NewInstruction(operation.Hash, 4),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestIndexExpressions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "[1, 2, 3][1 + 1]",
			expectedConstants: []interface{}{1, 2, 3, 1, 1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Array, 3),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Constant, 4),
				operation.NewInstruction(operation.Add),
				operation.NewInstruction(operation.Index),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input:             "{1: 2}[2 - 1]",
			expectedConstants: []interface{}{1, 2, 2, 1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Hash, 2),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Sub),
				operation.NewInstruction(operation.Index),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestFunctions(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `fn() { return 5 + 10 }`,
			expectedConstants: []interface{}{
				5,
				10,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0),
					operation.NewInstruction(operation.Constant, 1),
					operation.NewInstruction(operation.Add),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `fn() { 1; 2 }`,
			expectedConstants: []interface{}{
				1,
				2,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0),
					operation.NewInstruction(operation.Pop),
					operation.NewInstruction(operation.Constant, 1),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestFunctionsWithoutReturnValue(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `fn() { }`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.Return),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestCompilerScopes(t *testing.T) {
	compiler := New()
	if compiler.currentScope != 0 {
		t.Errorf("currentScope wrong. got=%d, want=%d", compiler.currentScope, 0)
	}

	globalSymbolTable := compiler.symbols
	compiler.emit(operation.Mul)
	compiler.enterScope()
	if compiler.currentScope != 1 {
		t.Errorf("currentScope wrong. got=%d, want=%d", compiler.currentScope, 1)
	}

	compiler.emit(operation.Sub)
	if len(compiler.scopes[compiler.currentScope].instructions) != 1 {
		t.Errorf("instructions length wrong. got=%d", len(compiler.scopes[compiler.currentScope].instructions))
	}

	last := compiler.scopes[compiler.currentScope].emitted
	if last.Opcode != operation.Sub {
		t.Errorf("emitted.Opcode wrong. got=%d, want=%d", last.Opcode, operation.Sub)
	}

	if compiler.symbols.Outer != globalSymbolTable {
		t.Errorf("compiler did not enclose symbolTable")
	}

	compiler.leaveScope()
	if compiler.currentScope != 0 {
		t.Errorf("currentScope wrong. got=%d, want=%d", compiler.currentScope, 0)
	}

	if compiler.symbols != globalSymbolTable {
		t.Errorf("compiler did not restore global symbol table")
	}

	if compiler.symbols.Outer != nil {
		t.Errorf("compiler modified global symbol table incorrectly")
	}

	compiler.emit(operation.Add)
	if len(compiler.scopes[compiler.currentScope].instructions) != 2 {
		t.Errorf("instructions length wrong. got=%d", len(compiler.scopes[compiler.currentScope].instructions))
	}

	last = compiler.scopes[compiler.currentScope].emitted
	if last.Opcode != operation.Add {
		t.Errorf("emitted.Opcode wrong. got=%d, want=%d", last.Opcode, operation.Add)
	}

	previous := compiler.scopes[compiler.currentScope].prevEmitted
	if previous.Opcode != operation.Mul {
		t.Errorf("prevEmitted.Opcode wrong. got=%d, want=%d", previous.Opcode, operation.Mul)
	}
}
func TestFunctionCalls(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `fn() { 26 }();`,
			expectedConstants: []interface{}{
				26,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0), // 26
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 1), // The function (compiled)
				operation.NewInstruction(operation.Call, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
	let noArg = fn() { 26 };
	noArg();
	`,
			expectedConstants: []interface{}{
				26,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0), // 26
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 1), // The function (compiled)
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Call, 0),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				let oneArg = fn(a) { };
				oneArg(24);
			`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.Return),
				},
				24,
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Call, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				let manyArg = fn(a, b, c) { };
				manyArg(24, 25, 26);
			`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.Return),
				},
				24,
				25,
				26,
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Call, 3),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				let oneArg = fn(a) { a };
				oneArg(24);
			`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.GetLocal, 0),
					operation.NewInstruction(operation.ReturnValue),
				},
				24,
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Call, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				let manyArg = fn(a, b, c) { a; b; c };
				manyArg(24, 25, 26);
			`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.GetLocal, 0),
					operation.NewInstruction(operation.Pop),
					operation.NewInstruction(operation.GetLocal, 1),
					operation.NewInstruction(operation.Pop),
					operation.NewInstruction(operation.GetLocal, 2),
					operation.NewInstruction(operation.ReturnValue),
				},
				24,
				25,
				26,
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.GetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Constant, 3),
				operation.NewInstruction(operation.Call, 3),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestLetStatementScopes(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
				let num = 55;
				fn() { num }
			`,
			expectedConstants: []interface{}{
				55,
				[]operation.Instruction{
					operation.NewInstruction(operation.GetGlobal, 0),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.SetGlobal, 0),
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				fn() {
					let num = 55;
					num
				}
			`,
			expectedConstants: []interface{}{
				55,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0),
					operation.NewInstruction(operation.SetLocal, 0),
					operation.NewInstruction(operation.GetLocal, 0),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 1),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `
				fn() {
					let a = 55;
					let b = 77;
					a + b
				}
			`,
			expectedConstants: []interface{}{
				55,
				77,
				[]operation.Instruction{
					operation.NewInstruction(operation.Constant, 0),
					operation.NewInstruction(operation.SetLocal, 0),
					operation.NewInstruction(operation.Constant, 1),
					operation.NewInstruction(operation.SetLocal, 1),
					operation.NewInstruction(operation.GetLocal, 0),
					operation.NewInstruction(operation.GetLocal, 1),
					operation.NewInstruction(operation.Add),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 2),
				operation.NewInstruction(operation.Pop),
			},
		},
	}
	runCompilerTests(t, tests)
}

func TestBuiltins(t *testing.T) {
	tests := []compilerTestCase{
		{
			input: `
				len([]);
				push([], 1);
			`,
			expectedConstants: []interface{}{1},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.GetBuiltin, 0),
				operation.NewInstruction(operation.Array, 0),
				operation.NewInstruction(operation.Call, 1),
				operation.NewInstruction(operation.Pop),
				operation.NewInstruction(operation.GetBuiltin, 4),
				operation.NewInstruction(operation.Array, 0),
				operation.NewInstruction(operation.Constant, 0),
				operation.NewInstruction(operation.Call, 2),
				operation.NewInstruction(operation.Pop),
			},
		},
		{
			input: `fn() { len([]) }`,
			expectedConstants: []interface{}{
				[]operation.Instruction{
					operation.NewInstruction(operation.GetBuiltin, 0),
					operation.NewInstruction(operation.Array, 0),
					operation.NewInstruction(operation.Call, 1),
					operation.NewInstruction(operation.ReturnValue),
				},
			},
			expectedInstructions: []operation.Instruction{
				operation.NewInstruction(operation.Constant, 0),
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

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
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
func testInstructions(expected []operation.Instruction, actual operation.Instruction) error {
	flattened := flattenInstructions(expected)
	if len(actual) != len(flattened) {
		return fmt.Errorf("wrong instructions length.\nwant=%q\ngot =%q", flattened, actual)
	}
	for i, ins := range flattened {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at %d.\nwant=%q\ngot =%q", i, flattened, actual)
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

		case []operation.Instruction:
			fn, ok := actual[i].(*data.CompiledFunction)
			if !ok {
				return fmt.Errorf("constant %d - not a function: %T",
					i, actual[i])
			}
			err := testInstructions(constant, fn.Instructions)
			if err != nil {
				return fmt.Errorf("constant %d - testInstructions failed: %s",
					i, err)
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
