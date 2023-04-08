package main

import (
	"fmt"
	"local-packages/list"
)

func main() {
	array := list.ArrayList{}
	linkedlist := list.LinkedList{}
	doublylinkedlist := list.DoublyLinkedList{}

	array.Init(10)
	linkedlist.Init()
	doublylinkedlist.Init()

	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	array.AddOnIndex(99, 0)
	fmt.Println(array.Size())
	array.RemoveOnIndex(0)
	/* array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0)
	array.RemoveOnIndex(0) */
	fmt.Println(array.Size())

	linkedlist.AddOnIndex(99, 0)
	fmt.Println(linkedlist.Size())

	doublylinkedlist.AddOnIndex(99, 0)
	fmt.Println(doublylinkedlist.Size())

}
