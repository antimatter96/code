package main

import "fmt"

func reverseIneger(x int) int {
	if x > 2147483647 || x < -2147483647 {
		return 0
	}

	i := 0

	for x != 0 {
		i *= 10
		i += x % 10
		x /= 10

		if i > 2147483647 || i < -2147483647 {
			return 0
		}
	}

	return i
}

func driver__reverseIneger() {
	fmt.Println(reverseIneger(123))
	fmt.Println(reverseIneger(-123))
	fmt.Println(reverseIneger(120))
	fmt.Println(reverseIneger(0))
	fmt.Println(reverseIneger(1534236469))
}
