package main

import "fmt"

func arraySign(nums []int) int {
	count := 1

	for _, e := range nums {
		if e == 0 {
			return 0
		}

		if e < 0 {
			count *= -1
		}
	}

	return count
}

func driver__arraySign() {
	test_cases := []struct {
		input  []int
		answer int
	}{
		{[]int{-1, -2, -3, -4, 3, 2, 1}, 1},
		{[]int{1, 5, 0, 2, -3}, 0},
		{[]int{-1, 1, -1, 1, -1}, -1},
	}

	for _, tc := range test_cases {
		fmt.Println(arraySign(tc.input) == tc.answer)
	}
}
