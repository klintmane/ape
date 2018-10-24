package eval

import "ape/interpreter/data"

func evalArray(elements []data.Data) *data.Array {
	return &data.Array{Elements: elements}
}
