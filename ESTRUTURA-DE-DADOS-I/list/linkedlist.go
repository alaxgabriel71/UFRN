package list

import (
	"errors"
)

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

// initializes the linked list with size 0
func (linkedlist *LinkedList) Init() {
	linkedlist.size = 0
}

// returns if the linked list is empty or not
func (linkedlist *LinkedList) isEmpty() bool {
	if linkedlist.size == 0 {
		return true
	} else {
		return false
	}
}

// verifies if the index range is valid
func (linkedlist *LinkedList) indexIsValid(index int) bool {
	if index >= 0 && index <= linkedlist.size {
		return true
	} else {
		return false
	}
}

// inserts a new value in the first linked list position
func (linkedlist *LinkedList) Add(value int) {
	actualNode := linkedlist.head
	newNode := Node{}
	if linkedlist.isEmpty() {
		linkedlist.head = &newNode
		newNode.value = value
		newNode.next = nil
		linkedlist.size++
	} else {
		for actualNode.next != nil {
			actualNode = actualNode.next
		}
		actualNode.next = &newNode
		newNode.value = value
		newNode.next = nil
		linkedlist.size++
	}
}

// inserts a new value at the indicated index
func (linkedlist *LinkedList) AddOnIndex(value int, index int) error {
	newNode := Node{}
	if index == 0 {
		if linkedlist.isEmpty() {
			linkedlist.Add(value)
		} else {
			newNode.next = linkedlist.head
			newNode.value = value
			linkedlist.head = &newNode
			linkedlist.size++
		}
		return nil
	} else if linkedlist.indexIsValid(index) {
		actualNode := linkedlist.head
		for i := 1; i <= linkedlist.size; i++ {
			if i == index {
				newNode.next = actualNode.next
				newNode.value = value
				actualNode.next = &newNode
				linkedlist.size++
			} else {
				actualNode = actualNode.next
			}
		}
		return nil
	} else {
		return errors.New("index out of range (AddOnIndex)")
	}
}

// removes the linked list's last item
func (linkedlist *LinkedList) Remove() error {
	if linkedlist.isEmpty() {
		return errors.New("empty list (Remove)")
	} else {
		linkedlist.size--
		return nil
	}
}

// removes the value indicated by the index from the list
func (linkedlist *LinkedList) RemoveOnIndex(index int) error {
	if linkedlist.isEmpty() {
		return errors.New("empty list (RemoveOnIndex)")
	} else if index == 0 {
		linkedlist.head = linkedlist.head.next
		linkedlist.size--
		return nil
	} else if index >= 1 && index < linkedlist.size {
		prevNode := linkedlist.head
		actualNode := prevNode.next
		nextNode := actualNode.next
		i := 1
		for index != i {
			prevNode = actualNode
			actualNode = actualNode.next
			nextNode = actualNode.next
			i++
		}
		prevNode.next = nextNode
		linkedlist.size--
		return nil
	} else {
		return errors.New("index out of range (RemoveOnIndex)")
	}
}

// returns the value at the indicated position
func (linkedlist *LinkedList) Get(index int) (int, error) {
	k := 0
	prevNode := Node{}
	actualNode := Node{}

	if index >= 0 && index < linkedlist.size {
		prevNode.next = linkedlist.head.next
		prevNode.value = linkedlist.head.value

		for k < index {
			if prevNode.next != nil {
				actualNode.next = prevNode.next.next
				actualNode.value = prevNode.next.value

				prevNode.next = actualNode.next
				prevNode.value = actualNode.value

				k++
			}
		}
		return prevNode.value, nil
	} else {
		return 0, errors.New("index out of range (Get)")
	}
}

// changes the value at the indicated index
func (linkedlist *LinkedList) Set(value int, index int) error {
	if index >= 0 && index < linkedlist.size {
		actualNode := linkedlist.head
		i := 0
		for index != i {
			actualNode = actualNode.next
			i++
		}
		actualNode.value = value
		return nil
	} else {
		return errors.New("index out of range (Set)")
	}
}

// returns the linked list's size
func (linkedlist *LinkedList) Size() int {
	return linkedlist.size
}
