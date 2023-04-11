package main

import (
	"fmt"
	"local-packages/list"
)

func main() {
	linkedlist := list.LinkedList{}

	linkedlist.Init()

	for i := 0; i < 10; i++ {
		linkedlist.Add(i)
	}
	//fmt.Println(linkedlist.Size())

	/* for i := 0; i < 10; i++ {
		val, _ := linkedlist.Get(i)
		fmt.Println(val)
	} */

	linkedlist.Reverse()

	for i := 0; i < 10; i++ {
		val, _ := linkedlist.Get(i)
		fmt.Println(val)
	}

	/* val, _ := linkedlist.Get(0)
	fmt.Println(val) */
}
