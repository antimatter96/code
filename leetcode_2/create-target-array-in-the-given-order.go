package main

func createTargetArray(nums []int, indexes []int) []int {
	n := len(nums)
	arr := make([]int, 0, n)

	for i := 0; i < n; i++ {
		index := indexes[i]

		if index > len(arr)/2 {
			y := make([]int, len(arr[index:]))
			copy(y, arr[index:])
			arr = append(append(arr[:index], nums[i]), y...)
		} else {
			x := make([]int, len(arr[:index]))
			copy(x, arr[:index])
			arr = append(append(x, nums[i]), arr[index:]...)
		}
	}

	return arr
}
