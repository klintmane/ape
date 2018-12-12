package vm

import "ape/src/data"

func (vm *VM) buildArray(startIndex, endIndex int) data.Data {
	elements := make([]data.Data, endIndex-startIndex)
	for i := startIndex; i < endIndex; i++ {
		elements[i-startIndex] = vm.stack.items[i]
	}
	return &data.Array{Elements: elements}
}
