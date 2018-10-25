package eval

import "ape/src/data"

func evalArray(elements []data.Data) *data.Array {
	return &data.Array{Elements: elements}
}
