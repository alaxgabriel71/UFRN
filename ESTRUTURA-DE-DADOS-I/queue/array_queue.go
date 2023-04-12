package queue

import (
	"local-packages/list"
)

type ArrayQueue struct {
	values list.ArrayList
	front  int
	rear   int
	size   int
}

func (arrayqueue *ArrayQueue) Init(capacity int) {
	arrayqueue.values.Init(capacity)
	arrayqueue.front = -1
	arrayqueue.rear = -1
	arrayqueue.size = 0
}

func (arrayqueue *ArrayQueue) Enqueue(value int) {
	if !arrayqueue.Empty() {
		arrayqueue.values.Add(value)
		arrayqueue.size++
		arrayqueue.front++
		arrayqueue.rear++
	} else {
		if arrayqueue.rear+1 != arrayqueue.front {
			arrayqueue.values.Add(value)
			arrayqueue.rear = (arrayqueue.rear + 1) % arrayqueue.size
		} else {
			// expand arrayqueue capacity
		}
	}
}

func (arrayqueue *ArrayQueue) Empty() bool {
	if arrayqueue.size > 0 {
		return false
	} else {
		return true
	}
}
