package main

import (
	"fmt"
	"reflect"
)

func nextGreaterElements(nums []int) []int {
	nums = append(nums, nums...)

	stk := intStack{}

	final := make([]int, len(nums))

	for i, e := range nums {
		if stk.Empty() {
			stk.Push(i)
			continue
		}

		for !stk.Empty() && e > nums[stk.Top()] {
			smallerIndex := stk.Pop()
			final[smallerIndex] = e
		}

		stk.Push(i)
	}

	for !stk.Empty() {
		smallerIndex := stk.Pop()
		final[smallerIndex] = -1
	}

	return final[0 : len(final)/2]
}

func driver__nextGreaterElements() {
	test_cases := []struct {
		input  []int
		output []int
	}{
		{
			[]int{1, 2, 1}, []int{2, -1, 2},
		}, {
			[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4},
		},
	}

	for _, tc := range test_cases {
		fmt.Println(reflect.DeepEqual(nextGreaterElements(tc.input), tc.output))
	}
}
