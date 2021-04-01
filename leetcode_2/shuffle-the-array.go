package main

func shuffle(nums []int, n int) []int {
	arr := make([]int, 2*n)
	for i := 0; i < n; i++ {
		arr[2*i] = nums[i]
	}
	for i := 0; i < n; i++ {
		arr[(2*i)+1] = nums[i+n]
	}
	return arr
}

func driver__shuffle() {

}
