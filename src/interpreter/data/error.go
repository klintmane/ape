package data

type Error struct {
	Message string
}

func (e *Error) Type() DataType  { return ERROR_TYPE }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
