package compiler

import (
	"ape/src/ast"
	"ape/src/compiler/operation"
	"ape/src/data"
	"fmt"
)

// Bytecode contains the instructions and constants the compiler generated and evaluated
type Bytecode struct {
	Instructions operation.Instruction
	Constants    []data.Data
}

// Compiler contains the instructions and constants which will then be turned into bytecode
type Compiler struct {
	instructions operation.Instruction
	constants    []data.Data
}

// New creates a new compiler
func New() *Compiler {
	return &Compiler{
		instructions: operation.Instruction{},
		constants:    []data.Data{},
	}
}

// Compile compiles an AST and populates the instructions and constants accordingly
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(operation.Pop)

	case *ast.InfixExpression:
		// Convert LessThan operations to GreaterThan ones
		if node.Operator == "<" {

			// First compile the right node, then the left node, unlike other infixes
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}

			err = c.Compile(node.Left)
			if err != nil {
				return err
			}

			c.emit(operation.GreaterThan)
			return nil
		}

		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(operation.Add)
		case "-":
			c.emit(operation.Sub)
		case "*":
			c.emit(operation.Mul)
		case "/":
			c.emit(operation.Div)
		case ">":
			c.emit(operation.GreaterThan)
		case "==":
			c.emit(operation.Equal)
		case "!=":
			c.emit(operation.NotEqual)
		default:
			return fmt.Errorf("unknown operator %s", node.Operator)
		}

	case *ast.IntegerLiteral:
		integer := &data.Integer{Value: node.Value}
		c.emit(operation.Constant, c.addConstant(integer))

	case *ast.Boolean:
		if node.Value {
			c.emit(operation.True)
		} else {
			c.emit(operation.False)
		}
	}
	return nil
}

// Bytecode produces bytecode out of the compiler result
func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// Adds a constant to the constant pool and returns its index so it can be referenced
func (c *Compiler) addConstant(d data.Data) int {
	c.constants = append(c.constants, d)
	return len(c.constants) - 1
}

// Adds an instruction to the instruction list and returns its index
func (c *Compiler) addInstruction(ins []byte) int {
	c.instructions = append(c.instructions, ins...)
	return len(c.instructions) - 1
}

// Generates a new instruction, adds it to the instruction list and returns the position
func (c *Compiler) emit(op operation.Opcode, operands ...int) int {
	ins := operation.NewInstruction(op, operands...)
	return c.addInstruction(ins)
}
