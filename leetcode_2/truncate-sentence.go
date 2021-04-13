package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	ans := &strings.Builder{}

	for _, r := range s {
		if k == 0 {
			break
		}

		if r == ' ' {
			k--
			if k == 0 {
				break
			}
		}

		ans.WriteRune(r)
	}

	return ans.String()
}

func driver__truncateSentence() {
	fmt.Println(truncateSentence("Hello how are you Contestant", 4))
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
	fmt.Println(truncateSentence("chopper is not a tanuki", 5))
}
