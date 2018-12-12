package compiler

import (
	"ape/src/ast"
	"ape/src/compiler/operation"
	"ape/src/compiler/symbols"
	"ape/src/data"
	"fmt"
	"sort"
)

// Bytecode contains the instructions and constants the compiler generated and evaluated
type Bytecode struct {
	Instructions operation.Instruction
	Constants    []data.Data
}

// Emitted represents an emitted instruction
type Emitted struct {
	Opcode   operation.Opcode
	Position int
}

// Compiler contains the instructions and constants which will then be turned into bytecode
type Compiler struct {
	instructions operation.Instruction
	constants    []data.Data
	symbols      *symbols.SymbolTable
	emitted      Emitted // The last emitted instruction
	prevEmitted  Emitted // The emitted instruction before that
}

// New creates a new compiler
func New() *Compiler {
	return &Compiler{
		instructions: operation.Instruction{},
		constants:    []data.Data{},
		symbols:      symbols.New(),
		emitted:      Emitted{},
		prevEmitted:  Emitted{},
	}
}

// NewWithState creates a new compiler
func NewWithState(s *symbols.SymbolTable, c []data.Data) *Compiler {
	return &Compiler{
		instructions: operation.Instruction{},
		constants:    c,
		symbols:      s,
		emitted:      Emitted{},
		prevEmitted:  Emitted{},
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

	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "!":
			c.emit(operation.Bang)
		case "-":
			c.emit(operation.Minus)
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

	case *ast.StringLiteral:
		str := &data.String{Value: node.Value}
		c.emit(operation.Constant, c.addConstant(str))

	case *ast.ArrayLiteral:
		for _, el := range node.Elements {
			err := c.Compile(el)
			if err != nil {
				return err
			}
		}
		c.emit(operation.Array, len(node.Elements))

	case *ast.HashLiteral:
		keys := []ast.Expression{}
		for k := range node.Pairs {
			keys = append(keys, k)
		}
		// Sort the keys as Go doesn't guarantee key/val order on iteration
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})
		for _, k := range keys {
			err := c.Compile(k)
			if err != nil {
				return err
			}
			err = c.Compile(node.Pairs[k])
			if err != nil {
				return err
			}
		}
		c.emit(operation.Hash, len(node.Pairs)*2)

	case *ast.IfExpression:
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		// Emit a `JumpNotTruthy` with a temporary operand
		jumpNotTruthyPos := c.emit(operation.JumpNotTruthy, 9999)
		err = c.Compile(node.Consequent)
		if err != nil {
			return err
		}
		if c.popEmitted() {
			c.preventPop()
		}

		// Emit a `Jump` with a temporary operand
		jumpPos := c.emit(operation.Jump, 9999)
		afterConsequentPos := len(c.instructions)
		c.changeOperand(jumpNotTruthyPos, afterConsequentPos)
		if node.Alternate == nil {
			c.emit(operation.Null)
		} else {
			err := c.Compile(node.Alternate)
			if err != nil {
				return err
			}
			if c.popEmitted() {
				c.preventPop()
			}
		}

		afterAlternatePos := len(c.instructions)
		c.changeOperand(jumpPos, afterAlternatePos)

	case *ast.LetStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		symbol := c.symbols.Define(node.Name.Value)
		c.emit(operation.SetGlobal, symbol.Index)

	case *ast.Identifier:
		symbol, ok := c.symbols.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("Variable %s is undefined", node.Value)
		}
		c.emit(operation.GetGlobal, symbol.Index)

	case *ast.BlockStatement:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
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
	pos := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return pos
}

// Generates a new instruction, adds it to the instruction list and returns the position
func (c *Compiler) emit(op operation.Opcode, operands ...int) int {
	ins := operation.NewInstruction(op, operands...)
	pos := c.addInstruction(ins)
	c.setEmitted(op, pos)
	return pos
}

// Updates the emitted and prevEmitted values
func (c *Compiler) setEmitted(op operation.Opcode, pos int) {
	c.prevEmitted = c.emitted
	c.emitted = Emitted{Opcode: op, Position: pos}
}

// Checks if the last emitted instruction is a pop instruction
func (c *Compiler) popEmitted() bool {
	return c.emitted.Opcode == operation.Pop
}

// Prevents a pop instruction from taking place by replacing it with the previous instruction
func (c *Compiler) preventPop() {
	c.instructions = c.instructions[:c.emitted.Position]
	c.emitted = c.prevEmitted
}

// Changes the operand of an instruction at the given position
func (c *Compiler) changeOperand(pos int, operand int) {
	opcode := operation.Opcode(c.instructions[pos])
	ins := operation.NewInstruction(opcode, operand)

	// replaces the instruction with the one created with the new operand
	for i := range ins {
		c.instructions[pos+i] = ins[i]
	}
}
