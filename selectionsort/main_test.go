package main

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tsc := []struct {
		desc   string
		input  []int
		output []int
	}{
		{
			desc:   "given [5, 3, 2, 6, 10] should return [2, 3, 5, 6, 10]",
			input:  []int{5, 3, 2, 6, 10},
			output: []int{2, 3, 5, 6, 10},
		},
	}

	for _, c := range tsc {
		t.Run(c.desc, func(t *testing.T) {
			output := selectionSort(c.input)

			for idx, val := range output {
				fmt.Printf("output: %v, expected: %v, index: %d\n", val, c.output[idx], idx)
			}
		})
	}
}
