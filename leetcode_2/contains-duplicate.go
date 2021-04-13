package main

import "fmt"

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, num := range nums {
		if _, ok := mp[num]; ok {
			return true
		}
		mp[num] = true
	}
	return false
}

func driver__containsDuplicate() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}) == true)
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}) == false)
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) == true)
}
