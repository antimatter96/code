package main

func numIdenticalPairs(nums []int) int {
	mp := make(map[int]int)
	for _, e := range nums {
		mp[e]++
	}

	count := 0
	for _, v := range mp {
		count += (v) * (v - 1)
	}

	return count / 2
}

func driver__numIdenticalPairs() {

}
