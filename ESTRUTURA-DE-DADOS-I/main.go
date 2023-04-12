package main

import (
	"fmt"
	"local-packages/stack"
)

func main() {
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

	llstack := stack.LinkedListStack{}

	llstack.Init()

	for i := 0; i < 10; i++ {
		llstack.Push(i)
	}

	for i := -1; i < 12; i++ {
		fmt.Println(llstack.Pop())
	}
}
