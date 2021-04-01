package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	unsorted := make([]int, len(nums))
	copy(unsorted, nums)
	sort.Ints(nums)

	mp := make(map[int]int)
	mp[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			mp[nums[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		unsorted[i] = mp[unsorted[i]]

	}

	return unsorted
}

func driver__smallerNumbersThanCurrent() {

}
