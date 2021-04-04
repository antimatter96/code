package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	ans := &strings.Builder{}

	ans.Grow(len(address) * 2)
	for _, r := range address {
		if r == '.' {
			ans.WriteString("[.]")
		} else {
			ans.WriteRune(r)
		}
	}

	return ans.String()
}

func driver__defangIPaddr() {
	fmt.Println(defangIPaddr("1.1.1.1") == "1[.]1[.]1[.]1")
	fmt.Println(defangIPaddr("255.100.50.0") == "255[.]100[.]50[.]0")
}
