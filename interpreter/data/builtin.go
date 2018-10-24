package data

type BuiltinFn func(args ...Data) Data

type Builtin struct {
	Fn BuiltinFn
}

func (b *Builtin) Type() DataType  { return BUILTIN_TYPE }
func (b *Builtin) Inspect() string { return "builtin function" }
