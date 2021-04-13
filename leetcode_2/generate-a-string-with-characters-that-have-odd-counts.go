package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	s := &strings.Builder{}
	if n%2 == 0 {
		s.WriteRune('a')
		n--
	}
	for i := 0; i < n; i++ {
		s.WriteRune('b')
	}

	return s.String()
}

func driver__generateTheString() {
	fmt.Println(generateTheString(4))
	fmt.Println(generateTheString(2))
	fmt.Println(generateTheString(7))
}
