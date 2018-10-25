package data

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() DataType { return INTEGER_TYPE }

func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

func (i *Integer) Hash() uint64 {
	return uint64(i.Value)
}
