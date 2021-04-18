package main

import (
	"fmt"
)

func checkIfPangram(sentence string) bool {
	mp := make(map[rune]bool)

	for _, r := range sentence {
		mp[r] = true

		if len(mp) == 26 {
			return true
		}
	}

	return len(mp) == 26
}

func driver__checkIfPangram() {
	test_cases := []struct {
		input  string
		output bool
	}{
		{"thequickbrownfoxjumpsoverthelazydog", true},
		{"leetcode", false},
	}

	for _, tc := range test_cases {
		fmt.Println(checkIfPangram(tc.input) == tc.output)
	}
}
