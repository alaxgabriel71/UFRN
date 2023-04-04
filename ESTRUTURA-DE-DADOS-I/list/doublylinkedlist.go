package list

type DoublyLinkedList struct {
	size int
	head *Node
	tail *Node
}

type Node struct {
	value int
	next *Node
	prev *Node
}

func (doublylinkedlist *DoublyLinkedList) Init() {
	doublylinkedlist.size = 0
}

func (doublylinkedlist *DoublyLinkedList) Add(value int) {
	newNode := Node{}
	newNode.value = value
	newNode.next = nil
	
	if doublylinkelist.size == 0 {
		doublylinkedlist.head = &newNode
		newNode.prev = nil
	} else {
		newNode.prev = doublylinkedlist.tail
		doublylinkedlist.tail.next = newNode
		doublylinkedlist.tail = &newNode
	}

	doublylinkedlist.size++
}
