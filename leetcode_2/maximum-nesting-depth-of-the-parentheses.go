package main

import "fmt"

func maxDepth(s string) int {
	count := 0
	max := 0
	for _, r := range s {
		if r == '(' {
			count++
			if count > max {
				max = count
			}
		} else if r == ')' {
			count--
		}
	}
	return max
}

func driver__maxDepth() {

	test_cases := []struct {
		input  string
		answer int
	}{
		{"(1+(2*3)+((8)/4))+1", 3},
		{"(1)+((2))+(((3)))", 3},
		{"1+(2*3)/(2-1)", 1},
		{"1", 0},
	}

	for _, tc := range test_cases {
		fmt.Println(maxDepth(tc.input), tc.answer)
	}
}
