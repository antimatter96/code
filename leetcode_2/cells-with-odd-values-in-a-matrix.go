package main

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
