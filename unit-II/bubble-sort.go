package main

import "fmt"

func BubbleSort(vec []int) {
	for i := 0; i < len(vec); i++ {
		for j := 0; j < len(vec)-i-1; j++ {
			if vec[j] > vec[j+1] {
				vec[j+1], vec[j] = vec[j], vec[j+1]
			}
			fmt.Println("i=",i,"j=",j,vec)
		}
	}
}

func main() {
	v := []int{1,2,3,5,5,6,7,5,9}
	fmt.Println("Before: ", v)
	BubbleSort(v)
	fmt.Println("After: ", v)
}