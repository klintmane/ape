package data

import "fmt"

type Closure struct {
	Fn   *CompiledFunction
	Free []Data
}

func (c *Closure) Type() DataType { return CLOSURE_TYPE }
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
