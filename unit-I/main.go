package main

import (
	"fmt"
	"local-packages/queue"
)

func main() {
	// arraystack teste
	/* arrstack := stack.ArrayStack{}

	arrstack.Init(10)
	fmt.Println(arrstack.Peek())

	for i := 0; i < 20; i++ {
		arrstack.Push(i)
	}

	for i := 0; i < 14; i++ {
		fmt.Println(arrstack.Pop())
	}

	fmt.Println(arrstack.Peek()) */

	
	// linkedliststack test
	/* llstack := stack.LinkedListStack{}

	llstack.Init()

	for i := 0; i < 10; i++ {
		llstack.Push(i)
	}

	for i := -1; i < 12; i++ {
		fmt.Println(llstack.Pop()) 
	}*/

	//arrayqueue test
	arrqueue := queue.ArrayQueue{}
	
	arrqueue.Init(10)

	for i := 0; i < 10; i++ {
		arrqueue.Enqueue(i)
		//fmt.Println(arrqueue.Size())
	}

	for i := 0; i < 10; i++ {
		val, _ := arrqueue.Dequeue()
		fmt.Println(val)
	}	

	/* arrqueue.Enqueue(0)
	arrqueue.Enqueue(1)
	val1, _:=arrqueue.Dequeue()
	val2, _:=arrqueue.Dequeue()
	fmt.Println(val1)
	fmt.Println(val2) */
}
