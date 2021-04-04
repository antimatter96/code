package main

import "fmt"

func numberOfSteps(num int) int {
	count := 0

	for num != 0 {
		count++
		if num%2 != 0 {
			num--
		} else {
			num /= 2
		}
	}

	return count
}

func driver__numberOfSteps() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
	fmt.Println(numberOfSteps(123))
}
