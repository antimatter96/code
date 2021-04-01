package main

import "fmt"

func mininumOfThree(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	if c <= a && c <= b {
		return c
	}

	return 1 << 16
}

func getVal(i, j int, matrix [][]int) int {
	if i < 0 || i >= len(matrix) {
		return 1 << 16
	}
	if j < 0 || j >= len(matrix[0]) {
		return 1 << 16
	}

	return matrix[i][j]
}

func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	max := 1 << 16

	minDist := make([][]int, m)
	for i := 0; i < m; i++ {
		minDist[i] = make([]int, n)

		for j := 0; j < n; j++ {
			minDist[i][j] = max
		}
	}

	for j := 0; j < n; j++ {
		minDist[n-1][j] = matrix[n-1][j]
	}

	for i := m - 2; i > -1; i-- {
		for j := 0; j < n; j++ {
			minDist[i][j] = matrix[i][j] + mininumOfThree(
				getVal(i+1, j-1, minDist),
				getVal(i+1, j, minDist),
				getVal(i+1, j+1, minDist),
			)
		}

	}

	min := 1 << 16
	for j := 0; j < n; j++ {
		if minDist[0][j] < min {
			min = minDist[0][j]
		}
	}

	return min
}

/*
func minFallingPathSum(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	var min int

	for i := m - 2; i > -1; i-- {
		for j := 0; j < n; j++ {
			min = math.MaxInt32
			if j > 0 {
				min = matrix[i+1][j-1]
			}
			if min > matrix[i+1][j] {
				min = matrix[i+1][j]
			}
			if j < n-1 {
				if min > matrix[i+1][j+1] {
					min = matrix[i+1][j+1]
				}
			}

			matrix[i][j] += min
		}
	}

	min = math.MaxInt32
	for j := 0; j < n; j++ {
		if matrix[0][j] < min {
			min = matrix[0][j]
		}
	}

	return min
}
*/

func driver__minFallingPathSum() {
	fmt.Println(minFallingPathSum([][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}))
	fmt.Println(minFallingPathSum([][]int{
		{-19, 57},
		{-40, -5},
	}))
	fmt.Println(minFallingPathSum([][]int{
		{-48},
	}))
}
