package main

import "fmt"

const balloon string = "balloon"
const double = "lo"

func maxNumberOfBalloons(text string) int {
	arr := make([]int, 26)

	for _, r := range text {
		arr[r-97]++
	}

	for _, r := range double {
		arr[r-97] /= 2
	}

	min := 10_0000

	for _, r := range balloon {
		if min > arr[r-97] {
			min = arr[r-97]
		}
	}

	return min
}

func driver__maxNumberOfBalloons() {
	test_cases := []struct {
		input  string
		output int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
	}

	for _, tc := range test_cases {
		fmt.Println(maxNumberOfBalloons(tc.input), tc.output)
	}
}
