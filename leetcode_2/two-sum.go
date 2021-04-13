package main

import "fmt"

func twoSum(nums []int, target int) []int {
	arr := make([]int, 2)
	mp := make(map[int]int)

	for i, num := range nums {
		if j, ok := mp[target-num]; ok {
			arr[0] = j
			arr[1] = i
			return arr
		}
		mp[num] = i
	}

	return arr
}

func driver__twoSum() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
