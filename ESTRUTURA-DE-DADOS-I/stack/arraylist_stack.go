package stack

import (
	"errors"
	"local-packages/list"
)

type ArrayStack struct {
	values list.ArrayList
	size   int
}

func (arraystack *ArrayStack) Init(capacity int) {
	arraystack.values.Init(capacity)
	arraystack.size = 0
}

func (arraystack *ArrayStack) Push(value int) {
	arraystack.values.Add(value)
	arraystack.size++
}

func (arraystack *ArrayStack) Pop() (int, error) {
	if !arraystack.Empty() {
		arraystack.size--
		return arraystack.values.Get(arraystack.size)
	} else {
		return -1, errors.New("empty stack")
	}
}

func (arraystack *ArrayStack) Peek() (int, error) {
	if !arraystack.Empty() {
		return arraystack.values.Get(arraystack.size - 1)
	} else {
		return -1, errors.New("empty stack")
	}
}

func (arraystack *ArrayStack) Empty() bool {
	if arraystack.size > 0 {
		return false
	} else {
		return true
	}
}

func (arraystack *ArrayStack) Size() int {
	return arraystack.size
}
