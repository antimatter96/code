package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
	incRow := make(map[int]int)
	incCol := make(map[int]int)

	for _, index := range indices {
		incRow[index[0]]++
		incCol[index[1]]++
	}

	count := 0
	var temp int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			temp = incRow[i] + incCol[j]
			if temp%2 == 1 {
				count++
			}
		}
	}

	return count
}

func driver__oddCells() {
	fmt.Println(oddCells(2, 3, [][]int{[]int{0, 1}, []int{1, 1}}))
	fmt.Println(oddCells(2, 2, [][]int{[]int{1, 1}, []int{0, 0}}))
}
