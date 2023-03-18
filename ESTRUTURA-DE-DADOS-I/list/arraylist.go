package list

type ArrayList struct {
	values []int
	size   int
}

// initialize the array's value with 10 of capacity
// and the array's size with 0
func (arraylist *ArrayList) Init() {
	arraylist.values = make([]int, 10)
	arraylist.size = 0
}

// returns the array's size
func (arraylist *ArrayList) Size() int {
	return arraylist.size
}

// duplicates the array's capacity if it's full
func (arraylist *ArrayList) checkCapacity() {
	if len(arraylist.values) == arraylist.size {
		newarray := make([]int, len(arraylist.values)*2)
		for i := 0; i < len(arraylist.values); i++ {
			newarray[i] = arraylist.values[i]
		}
		arraylist.values = newarray
	}
}

// validates the index range
func (arraylist *ArrayList) checkIndex(index int, add bool) bool {
	var limit int
	switch add {
	case true :
		limit = arraylist.size + 1
	default:
		limit = arraylist.size 
	}
	
	if index >=0 && index < limit {
		return true
	} else {
		return false
	}
}

// adds a value in the array's end
func (arraylist *ArrayList) Add(value int) {
	arraylist.checkCapacity()
	arraylist.values[arraylist.size] = value
	arraylist.size++
}

// adds a value to the array in the indicated index
func (arraylist *ArrayList) AddOnIndex(value int, index int) {
	if arraylist.checkIndex(index, true) {
		arraylist.checkCapacity()
		for i := arraylist.size; i > index; i-- {
			arraylist.values[i] = arraylist.values[i-1]
		}
		arraylist.values[index] = value
		arraylist.size++
	}
}

// removes the array's last value
func (arraylist *ArrayList) Remove() {
	if !arraylist.isEmpty() {
		arraylist.size--
	}
}

// removes the value on the indicated index
func (arraylist *ArrayList) RemoveOnIndex(index int) {
	if !arraylist.isEmpty() {
		if arraylist.checkIndex(index, false) {
			for i := index; i < arraylist.size; i++ {
				arraylist.values[i] = arraylist.values[i+1]
			}
			arraylist.size--
		}
	}
}