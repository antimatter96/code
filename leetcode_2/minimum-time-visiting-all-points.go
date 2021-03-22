package main

func minTimeToVisitAllPoints(points [][]int) int {
	abs := func(x int) int {
		if x > 0 {
			return x
		} else {
			return -x
		}
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	n := len(points)
	dist := 0
	for i := 1; i < n; i++ {
		dist += max(abs(points[i][0]-points[i-1][0]), abs(points[i][1]-points[i-1][1]))
	}

	return dist
}
