package stack

import (
	"errors"
	"local-packages/list"
)

type LinkedListStack struct {
	values list.LinkedList
	size   int
}

func (linkedliststack *LinkedListStack) Init() {
	linkedliststack.values.Init()
	linkedliststack.size = 0
}

func (linkedliststack *LinkedListStack) Push(value int) {
	linkedliststack.values.Add(value)
	linkedliststack.size++
}

func (linkedliststack *LinkedListStack) Pop() (int, error) {
	if !linkedliststack.Empty() {
		linkedliststack.size--
		return linkedliststack.values.Get(linkedliststack.size)
	} else {
		return -1, errors.New("empty stack")
	}
}

func (linkedliststack *LinkedListStack) Peek() (int, error) {
	if !linkedliststack.Empty() {
		return linkedliststack.values.Get(linkedliststack.size - 1)
	} else {
		return -1, errors.New("empty stack")
	}
}

func (linkedliststack *LinkedListStack) Empty() bool {
	if linkedliststack.size > 0 {
		return false
	} else {
		return true
	}
}

func (linkedliststack *LinkedListStack) Size() int {
	return linkedliststack.size
}
