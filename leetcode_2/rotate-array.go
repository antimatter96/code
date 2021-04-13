package main

import "fmt"

// func rotate(nums []int, k int) {
// 	ans := make([]int, len(nums))
// 	n := len(ans)
// 	var index int
// 	for i := 0; i < n; i++ {
// 		index = (i + k) % n
// 		ans[index] = nums[i]
// 	}
// 	nums = ans
// }

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	rotateSub(nums, 0, n-1)
	rotateSub(nums, 0, k-1)
	rotateSub(nums, k, n-1)
}

func rotateSub(arr []int, start, end int) {
	for i := 0; i < (end-start+1)/2; i++ {
		arr[start+i], arr[end-i] = arr[end-i], arr[start+i]
	}
}

// func rotateSub(arr []int, start, end int) {
// 	l := 0
// 	for i := start; i <= end && start+l < end-l; i++ {
// 		arr[start+l], arr[end-l] = arr[end-l], arr[start+l]
// 		l++
// 	}
// }

func driver__rotate() {
	l := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(l, 3)
	fmt.Println(l)

	l = []int{-1, -100, 3, 99}
	//fmt.Println(l)
	rotate(l, 2)
	fmt.Println(l)
	//fmt.Println(l)
}
