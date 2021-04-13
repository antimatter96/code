package main

import "fmt"

func findDuplicates(nums []int) []int {
	arr := make([]int, 0)

	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			index = nums[i] - 1
		} else {
			index = -nums[i] - 1
		}

		if nums[index] < 0 {
			arr = append(arr, index+1)
		} else {
			nums[index] = -nums[index]
			// will not work if elements are tripled
			// for tripes or more
			// move this outside the else block
		}
	}

	return arr
}

func driver__findDuplicates() {
	test_cases := []struct {
		input  []int
		output []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 2, 2}, []int{2}},
		{[]int{1, 2, 2, 2}, []int{}},
	}

	for _, tc := range test_cases {
		fmt.Println(findDuplicates(tc.input), tc.output)
	}
}
