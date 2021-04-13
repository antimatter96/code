package main

import "fmt"

func moveZeroes(nums []int) {
	zero := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if zero != -1 {
				nums[zero] = nums[i]
				nums[i] = 0
				i = zero
				zero = -1
			}
		} else if zero == -1 {
			zero = i
		}
	}
}

func moveZeroes4(nums []int) {
	l := 0
	var temp int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp = nums[i]
			nums[i] = nums[l]
			nums[l] = temp

			/*
				nums[i], nums[l] = nums[l], nums[i]
			*/

			l++
		}
	}
}

func driver__moveZeroes() {
	l := []int{0, 1, 0, 3, 12}
	moveZeroes4(l)
	fmt.Println(l)

	l = []int{0}
	moveZeroes4(l)
	fmt.Println(l)

	l = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	moveZeroes4(l)
	fmt.Println(l)
}
