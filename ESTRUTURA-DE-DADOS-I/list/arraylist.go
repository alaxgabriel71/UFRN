package list

import (
	"errors"
)

type ArrayList struct {
	values []int
	size   int
}

// initializes the array's value with 10 of capacity
// and the array's size with 0
func (arraylist *ArrayList) Init(capacity int) error {
	if capacity > 0 {
		arraylist.values = make([]int, capacity)
		arraylist.size = 0
		return nil
	} else {
		return errors.New("can't init arraylist with size <= 0")
	}
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
func (arraylist *ArrayList) checkIndex(index int) bool {
	if index >= 0 && index < arraylist.size {
		return true
	} else {
		return false
	}
}

// verifies if the array is empty
func (arraylist *ArrayList) isEmpty() bool {
	if arraylist.size == 0 {
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
func (arraylist *ArrayList) AddOnIndex(value int, index int) error {
	if arraylist.checkIndex(index) || index == arraylist.size {
		if index == arraylist.size {
			arraylist.Add(value)
		} else if arraylist.checkIndex(index) {
			arraylist.checkCapacity()
			for i := arraylist.size; i > index; i-- {
				arraylist.values[i] = arraylist.values[i-1]
			}
			arraylist.values[index] = value
			arraylist.size++
		}
		return nil
	} else {
		return errors.New("index out of range (AddOnIndex)")
	}
}

// removes the array's last value
func (arraylist *ArrayList) Remove() {
	if !arraylist.isEmpty() {
		arraylist.size--
	}
}

// removes the value on the indicated index
func (arraylist *ArrayList) RemoveOnIndex(index int) error {
	if !arraylist.isEmpty() {
		if arraylist.checkIndex(index) {
			for i := index; i < arraylist.size-1; i++ {
				arraylist.values[i] = arraylist.values[i+1]
			}
			arraylist.size--
		}
		return nil
	} else {
		return errors.New("index out of range (RemoveOnIndex)")
	}
}

// returns the array's value indicated by the index
func (arraylist *ArrayList) Get(index int) (int, error) {
	if arraylist.checkIndex(index) {
		var val int
		if !arraylist.isEmpty() {
			if arraylist.checkIndex(index) {
				val = arraylist.values[index]
			}
		}
		return val, nil
	} else {
		return -1, errors.New("index out of range (Get)")
	}
}

// replaces a value in the array's position indicated by the index
func (arraylist *ArrayList) Set(value int, index int) error {
	if arraylist.checkIndex(index) {
		if !arraylist.isEmpty() {
			if arraylist.checkIndex(index) {
				arraylist.values[index] = value
			}
		}
		return nil
	} else {
		return errors.New("index out of range (Set)")
	}
}

// returns the array's size
func (arraylist *ArrayList) Size() int {
	return arraylist.size
}
