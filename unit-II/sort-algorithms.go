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

func SelectionSort(vec []int) {
	var lower int
	for i:=0; i<len(vec)-1;i++ {
		lower = i
		for j:=i+1; j<len(vec);j++ {
			if vec[j] < vec[lower] {
				lower = j
			}
		}
		vec[i], vec[lower] = vec[lower], vec[i]
	}
}

func main() {
	v := []int{5,5,4,4,3,3,2,1}
	fmt.Println("Unsorted: \t", v)
	SelectionSort(v)
	fmt.Println("Sorted: \t", v)
}