package data

type String struct {
	Value string
}

func (s *String) Type() DataType  { return STRING_TYPE }
func (s *String) Inspect() string { return s.Value }
