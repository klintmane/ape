package data

type Return struct {
	Value Data
}

func (rv *Return) Type() DataType  { return RETURN_TYPE }
func (rv *Return) Inspect() string { return rv.Value.Inspect() }
