package main

import "fmt"

func permute(nums []int) [][]int {
	n := len(nums)
	initial := make([][]int, 1, n*n)
	initial[0] = make([]int, 0)

	for i := 0; i < n; i++ {
		n := len(initial)
		for j := 0; j < n; j++ {
			initial = append(initial, rec(initial[j], nums[i])...)
		}
		initial = initial[len(initial)-((i+1)*n):]
	}

	return initial
}

func rec(permutation []int, e int) [][]int {
	n := len(permutation)

	arr := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		arr[i] = make([]int, 0, n+1)
	}

	for i := 0; i < n; i++ {
		arr[i] = append(arr[i], permutation[0:i]...)
		arr[i] = append(arr[i], e)
		arr[i] = append(arr[i], permutation[i:n]...)
	}

	arr[n] = append(arr[n], permutation...)
	arr[n] = append(arr[n], e)

	return arr
}

func driver__permute() {
	test_cases := []struct {
		input  []int
		output [][]int
	}{
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3}, {1, 3, 2},
			{2, 1, 3}, {2, 3, 1},
			{3, 2, 1}, {3, 1, 2},
		}},
		{[]int{0, 1}, [][]int{
			{1, 0}, {0, 1},
		}},
		{[]int{1}, [][]int{
			{1},
		}},
	}

	for _, tc := range test_cases {
		fmt.Println(permute(tc.input), tc.output)
	}
}
