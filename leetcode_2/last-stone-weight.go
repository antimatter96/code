package main

import (
	"fmt"
)

func lastStoneWeight(stones []int) int {
	hp := arrayToBasicHeap(stones)

	for hp.Size() > 1 {
		x := hp.Remove()
		y := hp.Remove()

		if y != x {
			hp.Insert(x - y)
		}
	}

	var top int
	if hp.Size() > 0 {
		top = hp.Remove()
	}

	return top
}

func driver__lastStoneWeight() {
	test_cases := []struct {
		input  []int
		output int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{2, 2}, 0},
	}

	for _, tc := range test_cases {
		fmt.Println(lastStoneWeight(tc.input) == tc.output)
	}
}
