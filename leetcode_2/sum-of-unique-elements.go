package main

func sumOfUnique(nums []int) int {
	mp := make(map[int]int)

	for _, i := range nums {
		mp[i]++
	}

	sum := 0

	for k, v := range mp {
		if v == 1 {
			sum += k
		}
	}

	return sum
}

func driver__sumOfUnique() {

}
