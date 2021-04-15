package main

import "fmt"

func diStringMatch(s string) []int {
	arr := make([]int, 0, len(s)+1)

	max := len(s)
	min := 0

	for _, r := range s {
		if r == 'I' {
			arr = append(arr, min)
			min++
		} else {
			arr = append(arr, max)
			max--
		}
	}

	arr = append(arr, max)

	return arr
}

func diStringMatch_2(s string) []int {
	arr := make([]int, len(s)+1)

	max := len(s)
	min := 0
	index := 0

	for _, r := range s {
		if r == 'I' {
			arr[index] = min
			min++
		} else {
			arr[index] = max
			max--
		}
		index++
	}

	arr[index] = max

	return arr
}

func driver__diStringMatch() {
	test_cases := []struct {
		input  string
		output []int
	}{
		{"IDID", []int{0, 4, 1, 3, 2}},
		{"III", []int{0, 1, 2, 3}},
		{"DDI", []int{3, 2, 0, 1}},
	}

	for _, tc := range test_cases {
		fmt.Println(diStringMatch(tc.input), tc.output)
	}
}
