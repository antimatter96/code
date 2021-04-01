package main

import "fmt"

func minCostToMoveChips(position []int) int {
	odd := 0
	even := 0

	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > odd {
		return odd
	}
	return even
}

func driver__minCostToMoveChips() {
	fmt.Println(minCostToMoveChips([]int{1, 2, 3}))
	fmt.Println(minCostToMoveChips([]int{2, 2, 2, 3, 3}))
	fmt.Println(minCostToMoveChips([]int{1, 1000000}))
}
