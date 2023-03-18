package list

type ArrayList struct {
	values []int
	size   int
}

func (arraylist *ArrayList) Init() {
	arraylist.values = make([]int, 10)
	arraylist.size = 0
}

func (arraylist *ArrayList) Size() int {
	return arraylist.size
}
