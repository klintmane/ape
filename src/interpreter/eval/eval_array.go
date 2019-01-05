package eval

import "github.com/ape-lang/ape/src/data"

func evalArray(elements []data.Data) *data.Array {
	return &data.Array{Elements: elements}
}
