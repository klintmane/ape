package eval

import "github.com/ape-lang/ape/src/data"

func evalReturn(value data.Data) *data.Return {
	return &data.Return{Value: value}
}
