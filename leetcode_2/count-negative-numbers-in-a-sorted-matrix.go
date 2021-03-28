package main

func countNegatives(grid [][]int) int {
	binarySearchForFirstNegative := func(arr []int) int {
		if arr[0] < 0 {
			return 0
		}
		hi := len(arr) - 1
		lo := 0
		for lo <= hi && lo < len(arr) && hi > -1 {
			if arr[lo] < 0 {
				return lo
			}
			mid := (lo + hi) / 2
			if arr[mid] >= 0 {
				lo = mid + 1
			} else if arr[mid] <= 0 {
				hi = mid
			}
		}
		return -1
	}

	count := 0
	n := len(grid[0])
	for i := 0; i < len(grid); i++ {
		x := binarySearchForFirstNegative(grid[i])

		if x != -1 {
			count += (n - x)
		}
	}

	return count
}

func countNegatives2(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	count := 0

	i, j := 0, n-1
	for j > -1 && i < m {
		if grid[i][j] < 0 {
			count += m - i
			j--
		} else {
			i++
		}
	}

	return count
}
