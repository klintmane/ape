package eval

import "ape/src/interpreter/data"

func evalArray(elements []data.Data) *data.Array {
	return &data.Array{Elements: elements}
}
