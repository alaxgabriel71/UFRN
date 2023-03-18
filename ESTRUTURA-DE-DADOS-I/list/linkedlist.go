package list

type LinkedList struct {
	head *Node
	size int
}

type Node struct {
	value int
	next *Node
}

// initializes the linked list with size 0
func (linkedlist *LinkedList) Init() {
	linkedlist.size = 0
}

// returns the linked list's size
func (linkedlist *LinkedList) Size() int {
	return linkedlist.size
}