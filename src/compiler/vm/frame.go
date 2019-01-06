package vm

import (
	"github.com/ape-lang/ape/src/compiler/operation"
	"github.com/ape-lang/ape/src/data"
)

// Frame consists of a reference to a compiled function, an instruction pointer and a pointer to the base frame
type Frame struct {
	closure      *data.Closure
	pointer      int
	framePointer int
}

type Frames struct {
	items []*Frame
	index int
}

// NewFrame creates a new frame for a given function
func NewFrame(cl *data.Closure, framePointer int) *Frame {
	return &Frame{closure: cl, pointer: -1, framePointer: framePointer}
}

// NewFrames creates a collection of Frames
func NewFrames(max int) *Frames {
	return &Frames{
		items: make([]*Frame, max),
		index: 0,
	}
}

// Instructions returns the instructions of the function referenced byt the frame
func (f *Frame) Instructions() operation.Instruction {
	return f.closure.Fn.Instructions
}

// current frame in a frame collection
func (f *Frames) current() *Frame {
	return f.items[f.index-1]
}

// push a frame into a frame collection
func (f *Frames) push(frame *Frame) {
	f.items[f.index] = frame
	f.index++
}

// pop a frame from a frame collection
func (f *Frames) pop() *Frame {
	f.index--
	return f.items[f.index]
}
