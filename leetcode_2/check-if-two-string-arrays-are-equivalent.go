package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := strings.Join(word1, "")
	s2 := strings.Join(word2, "")

	return s1 == s2
}

func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	s1 := strings.Join(word1, "")
	s1Len := len(s1)

	idx := 0

	for _, w := range word2 {
		for i := 0; i < len(w); i++ {
			if idx >= s1Len || w[i] != s1[idx] {
				return false
			}
			idx++
		}
	}

	return idx == s1Len
}

func driver__arrayStringsAreEqual() {
	test_cases := []struct {
		w1, w2 []string
		output bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
		{[]string{"abc", "d", "defg"}, []string{"abcddefg"}, true},
		{[]string{"abc", "dddd", "defg"}, []string{"abcddefg", "aa"}, false},
	}

	for _, tc := range test_cases {
		fmt.Println(arrayStringsAreEqual(tc.w1, tc.w2), tc.output)
	}
}
