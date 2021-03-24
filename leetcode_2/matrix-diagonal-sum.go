package main

func diagonalSum(mat [][]int) int {
	n := len(mat)

	sum := 0
	if n%2 != 0 {
		sum -= mat[n/2][n/2]
	}

	for i := 0; i < n; i++ {
		sum += mat[i][i]
		sum += mat[i][n-i-1]
	}

	return sum

}
