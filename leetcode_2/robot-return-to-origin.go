package main

import "fmt"

func judgeCircle(moves string) bool {
	if len(moves)%2 != 0 {
		return false
	}

	x := 0
	y := 0

	for _, r := range moves {
		switch r {
		case 'L':
			x++
		case 'R':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}

	return x == 0 && y == 0
}

func driver__judgeCircle() {
	test_cases := []struct {
		input  string
		answer bool
	}{
		{"UD", true},
		{"LL", false},
		{"RRDD", false},
		{"LDRRLRUULR", false},
	}

	for _, tc := range test_cases {
		fmt.Println(judgeCircle(tc.input) == tc.answer)
	}
}
