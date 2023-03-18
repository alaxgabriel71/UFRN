package main

import (
	"fmt"
	"local-packages/list"
)

func main() {
	//array := list.ArrayList{}
	var array *list.LinkedList = new(list.LinkedList)

	array.Init()
	// array.Add(0)
	// array.Add(1)
	// array.Add(2)
	// array.Add(3)
	// array.AddOnIndex(99,1)
	// array.AddOnIndex(88,5)
	// array.AddOnIndex(77,7)
	// array.RemoveOnIndex(1)
	// array.RemoveOnIndex(5)
	// array.RemoveOnIndex(4)
	// array.RemoveOnIndex(0)
	// array.RemoveOnIndex(-1)
	// array.Set(55, 2)
	// array.Set(55, 3)
	// array.Set(55, -1)
	fmt.Println(array)
	fmt.Println("Size ->", array.Size())
	/* for i := 0; i < array.Size(); i++ {
		fmt.Println("[", i, "]=",array.Get(i))
	} */
}
