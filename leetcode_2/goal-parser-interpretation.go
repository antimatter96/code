package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {
	ans := &strings.Builder{}

	for i := 0; i < len(command); i++ {
		if command[i] == '(' {
			if command[i+1] == ')' {
				ans.WriteByte('o')
				i += 1
			} else {
				ans.WriteString("al")
				i += 3
			}
		} else {
			ans.WriteByte(command[i])
		}

	}

	//fmt.Println(ans.String())
	return ans.String()
}

func driver__interpret() {
	testCases := [][]string{
		{"G()(al)", "Goal"},
		{"G()()()()(al)", "Gooooal"},
		{"(al)G(al)()()G", "alGalooG"},
	}

	for _, tc := range testCases {
		fmt.Println(interpret(tc[0]) == tc[1])
	}
}
