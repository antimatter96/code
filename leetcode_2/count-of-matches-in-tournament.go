package main

import "fmt"

func numberOfMatches(n int) int {
	x := 0

	for n > 1 {
		x += n / 2
		n = (n % 2) + (n / 2)
	}

	return x
}

func driver__numberOfMatches() {
	fmt.Println(numberOfMatches(7))
	fmt.Println(numberOfMatches(14))
}
