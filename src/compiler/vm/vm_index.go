package vm

import (
	"fmt"

	"github.com/ape-lang/ape/src/data"
)

func (vm *VM) executeIndexExpr(left, index data.Data) error {
	switch {
	case left.Type() == data.ARRAY_TYPE && index.Type() == data.INTEGER_TYPE:
		return vm.executeArrayIndex(left, index)

	case left.Type() == data.HASH_TYPE:
		return vm.executeHashIndex(left, index)

	default:
		return fmt.Errorf("index operator not supported: %s", left.Type())
	}
}

func (vm *VM) executeArrayIndex(array, index data.Data) error {
	arr := array.(*data.Array)
	i := index.(*data.Integer).Value
	max := int64(len(arr.Elements) - 1)
	if i < 0 || i > max {
		return vm.stack.push(data.NULL)
	}
	return vm.stack.push(arr.Elements[i])
}

func (vm *VM) executeHashIndex(hash, index data.Data) error {
	hashObject := hash.(*data.Hash)
	key, ok := index.(data.HashableData)
	if !ok {
		return fmt.Errorf("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[data.HashData(key)]
	if !ok {
		return vm.stack.push(data.NULL)
	}
	return vm.stack.push(pair.Value)
}
