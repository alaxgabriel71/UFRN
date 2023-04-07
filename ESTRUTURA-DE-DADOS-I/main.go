package main

import (
	"fmt"
	"local-packages/list"
)

func main() {
	//array := list.ArrayList{}
	//var array *list.LinkedList = new(list.LinkedList)
	array := list.DoublyLinkedList{}

	array.Init()
	/* erro7 := array.RemoveOnIndex(0)
	if erro7 != nil {
		fmt.Println(erro7)
	} */
	array.Add(55)
	array.Add(66)
	array.Add(77)
	erro8 := array.AddOnIndex(88, 3)
	array.Set(99, 4)

	//erro6 := array.Set(99, 4)
	if erro8 != nil {
		fmt.Println(erro8)
	}

	/* if erro6 != nil {
		fmt.Println(erro6)
	} */
	/* erro9 := array.RemoveOnIndex(3)
	if erro9 != nil {
		fmt.Println(erro9)
	} */

	val0, err := array.Get(0)
	if err == nil {
		fmt.Println(val0)
	}

	val, err := array.Get(1)
	if err == nil {
		fmt.Println(val)
	}

	val3, err := array.Get(2)
	if err == nil {
		fmt.Println(val3)
	}

	val1, err1 := array.Get(3)
	if err1 == nil {
		fmt.Println(val1)
	} else {
		fmt.Println(err1)
	}

	// array.Add(0)
	// array.Add(1)
	// array.Add(2)
	// array.Add(3)
	//array.AddOnIndex(99, 1)
	// array.AddOnIndex(88,5)
	// array.AddOnIndex(77,7)
	// array.RemoveOnIndex(1)
	// array.RemoveOnIndex(5)
	// array.RemoveOnIndex(4)
	// array.RemoveOnIndex(0)
	// array.RemoveOnIndex(-1)

	// fmt.Println(array.Get(0))
	// fmt.Println(array.Get(1))
	// fmt.Println(array.Get(2))
	//array.RunReverse()
	fmt.Println(array)
	fmt.Println("Size ->", array.Size())
	//fmt.Println(array.Get(0))
	//fmt.Println(array.Get(7))
	/* for i := 0; i < array.Size(); i++ {
		fmt.Println("[", i, "]=",array.Get(i))
	} */

	/* var actualNode *list.Node = new(list.Node)
	var nextNode *list.Node = new(list.Node)

	actualNode = array.head
	nextNode = array.head.next

	for actualNode.next != nil {
		fmt.Println(actualNode.value)
		actualNode = nextNode.next
		nextNode = actualNode.next
	} */
}
