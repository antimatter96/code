package main

import "fmt"

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

func driver__uniqueOccurrences() {
	fmt.Println(uniqueOccurrences([]int{1, 2, 2, 1, 1, 3}))
	fmt.Println(uniqueOccurrences([]int{1, 2}))
	fmt.Println(uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}))
}
