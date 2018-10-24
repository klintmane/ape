package eval

import (
	"ape/src/ast"
	"ape/src/interpreter/data"
)

func evalHashLiteral(node *ast.HashLiteral, env *data.Environment) data.Data {
	pairs := make(map[data.HashKey]data.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(data.HashableData)
		if !ok {
			return evalError("unusable as hash key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := data.HashData(hashKey)
		pairs[hashed] = data.HashPair{Key: key, Value: value}
	}

	return &data.Hash{Pairs: pairs}
}
