package main

import "fmt"

func findNumbers(nums []int) int {
	noOfDigits := func(x int) int {
		i := 1

		for x > 9 {
			x /= 10
			i++
		}
		return i
	}

	count := 0

	for _, num := range nums {
		if noOfDigits(num)%2 == 0 {
			count++
		}
	}

	return count
}

func driver__findNumbers() {
	fmt.Println(findNumbers([]int{12, 345, 2, 6, 7896}))
	fmt.Println(findNumbers([]int{555, 901, 482, 1771, 0, 1, 10, 100}))
}
