package main

import "fmt"

func isAVowel(b byte) bool {
	return b == 65 || b == 69 || b == 73 || b == 79 || b == 85 || b == 97 || b == 101 || b == 105 || b == 111 || b == 117

}

func halvesAreAlike(s string) bool {
	count := 0

	half := len(s) / 2
	for i := 0; i < half; i++ {
		if isAVowel(s[i]) {
			count++
		}
		if isAVowel(s[i+half]) {
			count--
		}
	}

	return count == 0
}

func driver__halvesAreAlike() {
	fmt.Println(halvesAreAlike("book") == true)
	fmt.Println(halvesAreAlike("textbook") == false)
	fmt.Println(halvesAreAlike("MerryChristmas") == false)
	fmt.Println(halvesAreAlike("AbCdEfGh") == true)

}
