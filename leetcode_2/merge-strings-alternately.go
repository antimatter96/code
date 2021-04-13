package main

import (
	"fmt"
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0

	s := &strings.Builder{}
	s.Grow(len(word1) + len(word2))

	for ; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
		s.WriteByte(word1[i])
		s.WriteByte(word2[j])
	}

	for ; i < len(word1); i++ {
		s.WriteByte(word1[i])
	}

	for ; j < len(word2); j++ {
		s.WriteByte(word2[j])
	}

	return s.String()
}

func driver__mergeAlternately() {

	test_cases := []struct {
		word1, word2 string
		output       string
	}{
		{"abc", "pqr", "apbqcr"},
		{"ab", "pqrs", "apbqrs"},
		{"abcd", "pq", "apbqcd"},
	}

	for _, tc := range test_cases {
		fmt.Println(mergeAlternately(tc.word1, tc.word2) == tc.output)
	}

}
