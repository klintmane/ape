package data

import (
	"fmt"
)

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() DataType  { return BOOLEAN_TYPE }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
