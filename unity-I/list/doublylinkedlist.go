package list

import (
	"errors"
)

type DoublyLinkedList struct {
	size int
	head *DoublyNode
	tail *DoublyNode
}

type DoublyNode struct {
	value int
	next  *DoublyNode
	prev  *DoublyNode
}

// initializes the doubly linked list with size 0
func (doublylinkedlist *DoublyLinkedList) Init() {
	doublylinkedlist.size = 0
}

// verifies if the index range is valid
func (doublylinkedlist *DoublyLinkedList) indexIsValid(index int) bool {
	if index >= 0 && index < doublylinkedlist.size {
		return true
	} else {
		return false
	}
}

// prints the doubly linked list's values from end to beginning
/* func (doublylinkedlist *DoublyLinkedList) RunReverse() {
	actualNode := doublylinkedlist.tail
	for i := doublylinkedlist.size - 1; i >= 0; i-- {
		fmt.Println(actualNode.value)
		actualNode = actualNode.prev
	}
} */

// inserts a new value in the first doubly linked list position
func (doublylinkedlist *DoublyLinkedList) Add(value int) {
	newNode := DoublyNode{}
	newNode.value = value
	newNode.next = nil

	if doublylinkedlist.size == 0 {
		doublylinkedlist.head = &newNode
		newNode.prev = nil
	} else {
		doublylinkedlist.tail.next = &newNode
		newNode.prev = doublylinkedlist.tail
	}

	doublylinkedlist.tail = &newNode
	doublylinkedlist.size++
}

// inserts a new value at the indicated index
func (doublylinkedlist *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index == 0 {
		if doublylinkedlist.size == 0 {
			doublylinkedlist.Add(value)
		} else {
			newNode := DoublyNode{}
			newNode.value = value
			newNode.prev = nil
			newNode.next = doublylinkedlist.head
			doublylinkedlist.head.prev = &newNode
			doublylinkedlist.head = &newNode
			doublylinkedlist.size++
		}
		return nil
	} else if index == doublylinkedlist.size {
		doublylinkedlist.Add(value)
		return nil
	} else if doublylinkedlist.indexIsValid(index) {
		actualNode := doublylinkedlist.head
		newNode := DoublyNode{}
		newNode.value = value
		i := 1
		for i != index {
			actualNode = actualNode.next
			i++
		}
		newNode.prev = actualNode
		newNode.next = actualNode.next
		actualNode.next.prev = &newNode
		actualNode.next = &newNode
		doublylinkedlist.size++
		return nil
	} else {
		return errors.New("index out of range (AddOnIndex)")
	}
}

// removes the doubly linked list's last item
func (doublylinkedlist *DoublyLinkedList) Remove() error {
	if doublylinkedlist.size == 0 {
		return errors.New("empty list")
	} else {
		if doublylinkedlist.size == 1 {
			doublylinkedlist.head = nil
			doublylinkedlist.tail = nil
			doublylinkedlist.size--
		} else {
			doublylinkedlist.tail = doublylinkedlist.tail.prev
			doublylinkedlist.tail.next = nil
			doublylinkedlist.size--
		}
		return nil
	}
}

// removes the value indicated by the index from the list
func (doublylinkedlist *DoublyLinkedList) RemoveOnIndex(index int) error {
	if doublylinkedlist.indexIsValid(index) {
		if index == doublylinkedlist.size-1 {
			doublylinkedlist.Remove()
		} else if index == 0 {
			doublylinkedlist.head = doublylinkedlist.head.next
			doublylinkedlist.head.prev = nil
			doublylinkedlist.size--
		} else {
			actualNode := doublylinkedlist.head
			i := 0
			for i != index {
				actualNode = actualNode.next
				i++
			}
			actualNode.prev.next = actualNode.next
			actualNode.next.prev = actualNode.prev
			doublylinkedlist.size--
		}
		return nil
	} else {
		return errors.New("index out of range (RemoveOnIndex)")
	}
}

// returns the value at the indicated position
func (doublylinkedlist *DoublyLinkedList) Get(index int) (int, error) {
	if doublylinkedlist.indexIsValid(index) {
		actualNode := doublylinkedlist.head
		for i := 0; i != index; i++ {
			actualNode = actualNode.next
		}
		return actualNode.value, nil
	} else {
		return -1, errors.New("index out of range (Get)")
	}

}

// changes the value at the indicated index
func (doublylinkedlist *DoublyLinkedList) Set(value int, index int) error {
	if doublylinkedlist.indexIsValid(index) {
		if index == doublylinkedlist.size-1 {
			doublylinkedlist.tail.value = value
		} else {
			actualNode := doublylinkedlist.head
			i := 0
			for i != index {
				actualNode = actualNode.next
				i++
			}
			actualNode.value = value
		}
		return nil
	} else {
		return errors.New("index out of range (Set)")
	}
}

// returns the doubly linked list's size
func (doublylinkedlist *DoublyLinkedList) Size() int {
	return doublylinkedlist.size
}
