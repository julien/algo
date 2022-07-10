package main

import (
	"testing"
)

func TestFact(t *testing.T) {
	tsc := []struct {
		desc   string
		input  int
		output int
	}{
		{
			desc:   "given fact(3) should return 6",
			input:  3,
			output: 6,
		},
		{
			desc:   "given fact(2) should return 2",
			input:  2,
			output: 2,
		},
		{
			desc:   "given fact(1) should return 1",
			input:  1,
			output: 1,
		},
		{
			desc:   "given fact(5) should return 120",
			input:  5,
			output: 120,
		},
	}

	for _, c := range tsc {
		t.Run(c.desc, func(t *testing.T) {
			output := fact(c.input)

			if output != c.output {
				t.Fatalf("want: %v, got: %v\n", c.output, output)
			}
		})
	}
}
