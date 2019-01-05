package eval

import "github.com/ape-lang/ape/src/data"

func evalInteger(value int64) *data.Integer {
	return &data.Integer{Value: value}
}
