package main

import "testing"

func TestBinarySearch(t *testing.T) {
	list := []string{"a", "b", "c"}

	cases := []struct {
		name   string
		search string
		result int
	}{
		{
			name:   "gets the first index",
			search: "a",
			result: 0,
		},
		{
			name:   "gets the last index",
			search: "c",
			result: 2,
		},
		{
			name:   "fails getting a non existing element",
			search: "Z",
			result: -1,
		},
		{
			name:   "fails getting a non existing element",
			search: "A",
			result: -1,
		},
		{
			name:   "fails getting a non existing element",
			search: "d",
			result: -1,
		},
	}

	for _, c := range cases {
		result := binary_search(list, c.search)
		if result != c.result {
			t.Errorf("expected %s, got %v\n", c.search, result)
		}
	}

}
