package eval

import "ape/src/data"

func evalReturn(value data.Data) *data.Return {
	return &data.Return{Value: value}
}
