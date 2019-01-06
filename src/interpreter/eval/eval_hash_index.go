package eval

import "github.com/ape-lang/ape/src/data"

func evalHashIndexExpression(hash, index data.Data) data.Data {
	hashData := hash.(*data.Hash)
	key, ok := index.(data.HashableData)

	if !ok {
		return evalError("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashData.Pairs[data.HashData(key)]
	if !ok {
		return data.NULL
	}

	return pair.Value
}
