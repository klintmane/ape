package data

import (
	"fmt"

	"github.com/ape-lang/ape/src/compiler/operation"
)

type CompiledFunction struct {
	Instructions operation.Instruction
	LocalCount   int // local variable definitions count (needed to allocate space for the function in the stack)
	ParamCount   int // parameter count
}

func (cf *CompiledFunction) Type() DataType { return COMPILED_FUNCTION_TYPE }

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
