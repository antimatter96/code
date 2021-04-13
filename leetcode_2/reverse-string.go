package main

import "fmt"

func reverseString(s []byte) {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func reverseString2(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func driver__reverseString() {
	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s1)
	reverseString(s1)
	fmt.Println(string(s1))

	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	fmt.Println(s2)
	reverseString(s2)
	fmt.Println(string(s2))
}
