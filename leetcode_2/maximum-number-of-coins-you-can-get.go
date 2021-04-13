package main

import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	// sort.Slice is faster than sort.Ints
	sort.Slice(piles, func(i, j int) bool { return piles[i] < piles[j] })
	//sort.Ints(piles)

	n := len(piles) / 3
	sum := 0

	for i := len(piles) - 2; n > 0 && i > -1; i -= 2 {
		sum += piles[i]
		n--
	}

	return sum
}

func driver__maxCoins() {
	test_cases := []struct {
		piles  []int
		output int
	}{
		{[]int{2, 4, 1, 2, 7, 8}, 9},
		{[]int{2, 4, 5}, 4},
		{[]int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
	}

	for _, tc := range test_cases {
		fmt.Println(maxCoins(tc.piles), tc.output)
	}
}
