package eval

import "ape/interpreter/data"

func evalReturn(value data.Data) *data.Return {
	return &data.Return{Value: value}
}
