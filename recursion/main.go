package main

import "fmt"

func main() {
	fmt.Println("countdown: ")
	countdown(4)
	fmt.Println("---------")

	fmt.Println(fact(3))
}

func countdown(i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	} else {
		countdown(i - 1)
	}
}

func fact(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * fact(x-1)
	}
}
