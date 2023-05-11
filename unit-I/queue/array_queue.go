package queue

import (
	"errors"
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
	if arrayqueue.Empty() {
		arrayqueue.values.Add(value)
		arrayqueue.front++
		arrayqueue.rear++
		arrayqueue.size++
	} else {
		if ((arrayqueue.rear+1)%arrayqueue.size+1) != arrayqueue.front {
			arrayqueue.values.Add(value)
			arrayqueue.size++
			arrayqueue.rear = (arrayqueue.rear + 1) % arrayqueue.size
		} else {
			// expand arrayqueue capacity
			newArray := list.ArrayList{}
			newArray.Init(arrayqueue.size*2)

			for i := 0; i < arrayqueue.size; i++ {
				val, _ := arrayqueue.Dequeue()
				newArray.Add(val)
			}

			/* arrayqueue.values = newArray
			arrayqueue.values.Add(value)
			arrayqueue.front = 0
			arrayqueue.rear = arrayqueue.size
			arrayqueue.size++ */
		}
	}
	
}

func (arrayqueue *ArrayQueue) Dequeue() (int, error) {
	if arrayqueue.Empty() {
		return -1, errors.New("empty queue")
	} else {
		if ((arrayqueue.front+1)%arrayqueue.size) != arrayqueue.rear {
			val, _ := arrayqueue.values.Get(arrayqueue.front)
			/* fmt.Println("size",arrayqueue.size)
			fmt.Println("front", arrayqueue.front) */
			arrayqueue.front = (arrayqueue.front+1)%arrayqueue.size
			arrayqueue.size--
			return val, nil
		} else {
			val, _ := arrayqueue.values.Get(arrayqueue.front)
			arrayqueue.size--
			arrayqueue.front = -1
			arrayqueue.rear = -1
			return val, nil
		}
	}
}

func (arrayqueue *ArrayQueue) Front() (int, error) {
	if arrayqueue.Empty() {
		return -1, errors.New("empty queue")
	} else {
		return arrayqueue.values.Get(arrayqueue.front)
	}
}

func (arrayqueue *ArrayQueue) Empty() bool {
	if arrayqueue.size > 0 {
		return false
	} else {
		return true
	}
}

func (arrayqueue *ArrayQueue) Size() int {
	return arrayqueue.size
}
