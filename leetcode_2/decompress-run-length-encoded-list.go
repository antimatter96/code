package main

func decompressRLElist(nums []int) []int {
	total := 0
	for i := 0; i < len(nums); i += 2 {
		total += nums[i]
	}

	arr := make([]int, 0, total)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			arr = append(arr, nums[i+1])
		}
	}

	return arr
}

func driver__decompressRLElist() {

}
