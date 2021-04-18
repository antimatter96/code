package main

import (
	"fmt"
	"reflect"
)

func finalPrices(nums []int) []int {
	stk := intStack{}

	dicount := make([]int, len(nums))

	for i, e := range nums {
		if stk.Empty() {
			stk.Push(i)
			continue
		}

		for !stk.Empty() && e <= nums[stk.Top()] {
			smallerIndex := stk.Pop()
			dicount[smallerIndex] = e
		}

		stk.Push(i)
	}

	// for !stk.Empty() {
	// 	smallerIndex := stk.Pop()
	// 	dicount[smallerIndex] = 0
	// }

	for i := 0; i < len(nums); i++ {
		nums[i] -= dicount[i]
	}
	return nums
}

func driver__finalPrices() {
	test_cases := []struct {
		input  []int
		output []int
	}{
		{
			[]int{8, 4, 6, 2, 3}, []int{4, 2, 4, 2, 3},
		}, {
			[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5},
		},
		{
			[]int{10, 1, 1, 6}, []int{9, 0, 1, 6},
		},
	}

	for _, tc := range test_cases {
		fmt.Println(reflect.DeepEqual(finalPrices(tc.input), tc.output))
	}
}
