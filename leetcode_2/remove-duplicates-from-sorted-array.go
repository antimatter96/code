package main

import "fmt"

func removeDuplicatesFromSortedarray(nums []int) int {
	writePointer := 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[writePointer-1] {

		} else {
			nums[writePointer] = nums[i]
			writePointer++
		}
	}
	return writePointer
}

func driver__removeDuplicatesFromSortedarray() {
	fmt.Println(removeDuplicatesFromSortedarray([]int{1, 1, 2}), 2)
	fmt.Println(removeDuplicatesFromSortedarray([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}), 5)

}
