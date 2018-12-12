package vm

import (
	"ape/src/data"
	"fmt"
)

func (vm *VM) buildHash(startIndex, endIndex int) (data.Data, error) {
	hashedPairs := make(map[data.HashKey]data.HashPair)
	for i := startIndex; i < endIndex; i += 2 {
		key := vm.stack.items[i]
		value := vm.stack.items[i+1]
		pair := data.HashPair{Key: key, Value: value}
		hashKey, ok := key.(data.HashableData)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", key.Type())
		}
		hashedPairs[data.HashData(hashKey)] = pair
	}
	return &data.Hash{Pairs: hashedPairs}, nil
}
