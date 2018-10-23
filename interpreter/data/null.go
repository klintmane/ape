package data

type Null struct{}

func (n *Null) Type() DataType  { return NULL_TYPE }
func (n *Null) Inspect() string { return "null" }
