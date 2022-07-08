package main

import "fmt"

func main() {
	a := []int{5, 3, 6, 2, 10}
	fmt.Printf("input:\t%v\n", a)
	fmt.Printf("output:\t%v\n", selectionSort(a))
}

func findSmallest(src []int) int {
	sm := src[0]
	idx := 0
	for i := 1; i < len(src); i++ {
		if src[i] < sm {
			sm = src[i]
			idx = i
		}
	}
	return idx
}

func selectionSort(src []int) []int {
	size := len(src)
	out := make([]int, size)
	for i := 0; i < size; i++ {
		sm := findSmallest(src)
		out[i] = src[sm]
		src = append(src[:sm], src[sm+1:]...)
	}
	return out
}
