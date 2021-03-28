package main

func uniqueOccurrences(arr []int) bool {
	mp1 := make(map[int]int)
	for _, v := range arr {
		mp1[v]++
	}
	mp2 := make(map[int]int)
	for _, v := range mp1 {
		mp2[v]++
		if mp2[v] > 1 {
			return false
		}
	}

	return true
}
