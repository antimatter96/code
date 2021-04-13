package main

import "fmt"

func singleNumber(nums []int) int {
	x := 0

	for i := 0; i < len(nums); i++ {
		x ^= nums[i]
	}

	return x
}

func driver__singleNumber() {
	fmt.Println(singleNumber([]int{2, 2, 1}) == 1)
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}) == 4)
	fmt.Println(singleNumber([]int{1}) == 1)
}
