package eval

import "ape/src/data"

func evalInteger(value int64) *data.Integer {
	return &data.Integer{Value: value}
}