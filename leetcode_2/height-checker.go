package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	arr := make([]int, len(heights))

	copy(arr, heights)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	count := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != heights[i] {
			count++
		}
	}

	return count
}

func driver__heightChecker() {
	test_cases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}

	for _, tc := range test_cases {
		fmt.Println(heightChecker(tc.input), tc.output)
	}
}
