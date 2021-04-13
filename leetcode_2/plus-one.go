package main

import "fmt"

func plusOne(digits []int) []int {
	carry := 1

	for i := len(digits) - 1; i > -1; i-- {
		digits[i] += carry
		if digits[i] >= 10 {
			carry = 1
			digits[i] = 0
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func driver__plusOne() {
	fmt.Println(plusOne([]int{1, 2, 3}), []int{1, 2, 4})
	fmt.Println(plusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})
	fmt.Println(plusOne([]int{0}), []int{1})
	fmt.Println(plusOne([]int{9}), []int{1, 0})
	fmt.Println(plusOne([]int{9, 9}), []int{1, 0, 0})
	fmt.Println(plusOne([]int{2, 9}), []int{3, 0})
}
