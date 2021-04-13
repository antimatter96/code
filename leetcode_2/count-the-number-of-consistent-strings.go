package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	mp := make(map[rune]bool)
	for _, r := range allowed {
		mp[r] = true
	}

	count := len(words)

	for _, word := range words {
		for _, r := range word {
			if _, ok := mp[r]; !ok {
				count--
				break
			}
		}
	}

	return count
}

func countConsistentStrings2(allowed string, words []string) int {
	mp := make([]bool, 26)
	for _, r := range allowed {
		mp[r-'a'] = true
	}

	count := 0

	for _, word := range words {
		match := true
		for _, r := range word {
			if !mp[r-'a'] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}

	return count
}

func driver__countConsistentStrings() {
	fmt.Println(countConsistentStrings2("ab", []string{"ad", "bd", "aaab", "baa", "badab"}))

	fmt.Println(countConsistentStrings2("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}))

	fmt.Println(countConsistentStrings2("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}))

}
