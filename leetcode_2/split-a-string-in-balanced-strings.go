package main

import "fmt"

func balancedStringSplit(s string) int {
	counter := 0
	k := 0

	for _, r := range s {
		if r == 'R' {
			counter++
		} else {
			counter--
		}

		if counter == 0 {
			k++
		}
	}
	return k
}

func driver__balancedStringSplit() {
	test_cases := []struct {
		input  string
		answer int
	}{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}

	for _, tc := range test_cases {
		fmt.Println(balancedStringSplit(tc.input), tc.answer)
	}
}
