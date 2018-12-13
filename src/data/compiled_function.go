package data

import (
	"ape/src/compiler/operation"
	"fmt"
)

type CompiledFunction struct {
	Instructions operation.Instruction
}

func (cf *CompiledFunction) Type() DataType { return COMPILED_FUNCTION_TYPE }

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
