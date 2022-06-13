package main

import "fmt"

func main() {
	list := []string{"a", "b", "c", "d", "e", "f", "g", "h",
		"i", "j", "k", "l", "m", "n", "o", "p", "q",
		"r", "s", "t", "u", "v", "w", "x", "z"}

	fmt.Printf("searching for \"j\": %v\n", binary_search(list, "j"))
	fmt.Printf("searching for \"A\": %v\n", binary_search(list, "A"))
}

func binary_search(list []string, item string) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high)
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// With generics (requires go >= 1.18)
// import "golang.org/x/exp/constraints"

// func binary_search[T constraints.Ordered](list []T, item T) int {
// 	low := 0
// 	high := len(list) - 1

// 	for low <= high {
// 		mid := (low + high)
// 		guess := list[mid]

// 		if guess == item {
// 			return mid
// 		}
// 		if guess > item {
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}
// 	return -1
// }
