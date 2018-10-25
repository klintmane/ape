package eval

import "ape/src/data"

func evalArrayIndexExpression(array, index data.Data) data.Data {
	arrayData := array.(*data.Array)
	i := index.(*data.Integer).Value
	max := int64(len(arrayData.Elements) - 1)

	if i < 0 || i > max {
		return NULL
	}

	return arrayData.Elements[i]
}