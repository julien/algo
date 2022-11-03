package main

import "fmt"

func main() {
	fmt.Println(sum([]int{2, 4, 6}))
	fmt.Println(count([]int{5, 5, 4, 6, 2, 3}))
	fmt.Println(max([]int{5, 25, 42, 160, 12, 93}))
	fmt.Println(quicksort([]int{5, 25, 142, 160, 312, 93}))
}

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + sum(a[1:])
}

func count(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return 1 + count(a[1:])
}

func max(a []int) int {
	if len(a) == 2 {
		if a[0] > a[1] {
			return a[0]
		}
		return a[1]
	}

	submax := max(a[1:])
	if a[0] > submax {
		return a[0]
	}
	return submax
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	pivot := a[0]
	var smaller = []int{}
	var greater = []int{}
	for _, n := range a[1:] {
		if pivot > n {
			smaller = append(smaller, n)
		} else {
			greater = append(greater, n)
		}
	}

	smaller = append(quicksort(smaller), pivot)
	greater = quicksort(greater)
	return append(smaller, greater...)

}
