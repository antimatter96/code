package main

import "fmt"

func subsets(nums []int) [][]int {
	n := 1 << len(nums)
	arr := make([][]int, n)

	var z int
	for digit := 0; digit < len(nums); digit++ {
		z = 1 << digit
		for i := 0; i < n; i++ {
			if (i & z) > 0 {
				arr[i] = append(arr[i], nums[digit])
			}
		}
	}

	return arr
}

func drive__subsets() {
	test_cases := []struct {
		nums   []int
		output [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
			}},
		{[]int{0}, [][]int{
			{}, {0},
		}},
	}

	for _, tc := range test_cases {
		fmt.Println(subsets(tc.nums), tc.output, len(tc.output))
	}
}
