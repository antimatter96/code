package main

import (
	"fmt"
	"strings"
)

var codes = []string{
	".-",
	"-...",
	"-.-.",
	"-..",
	".",
	"..-.",
	"--.",
	"....",
	"..",
	".---",
	"-.-",
	".-..",
	"--",
	"-.",
	"---",
	".--.",
	"--.-",
	".-.",
	"...",
	"-",
	"..-",
	"...-",
	".--",
	"-..-",
	"-.--",
	"--..",
}

func mosrse(word string) string {
	//count := len(s)
	ans := &strings.Builder{}
	ans.Grow(3 * len(word))

	for _, r := range word {
		ans.WriteString(codes[r-97])
	}

	return ans.String()
}

func uniqueMorseRepresentations(words []string) int {
	mp := make(map[string]int)
	for _, word := range words {
		mp[mosrse(word)]++
	}

	return len(mp)
}

func driver__uniqueMorseRepresentations() {
	fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
