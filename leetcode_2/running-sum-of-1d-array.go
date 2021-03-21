package main

func runningSum(nums []int) []int {
	sum := 0
	arr := make([]int, len(nums))
	for i, e := range nums {
		sum += e
		arr[i] = sum
	}
	return arr
}
