package main

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)

	max := 0
	product := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			product = (nums[i] - 1) * (nums[j] - 1)
			if product > max {
				max = product
			}
		}
	}

	return max
}

func driver__maxProduct() {
	fmt.Println(maxProduct([]int{3, 4, 5, 2}))
	fmt.Println(maxProduct([]int{1, 5, 4, 5}))
	fmt.Println(maxProduct([]int{3, 7}))
}
