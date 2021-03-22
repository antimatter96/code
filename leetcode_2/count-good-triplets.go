package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	n := len(arr)

	abs := func(x int) int {
		if x > 0 {
			return x
		} else {
			return -x
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < n; k++ {
				if abs(arr[j]-arr[k]) > b {
					continue
				}

				if abs(arr[i]-arr[k]) > c {
					continue
				}

				count++
			}
		}
	}

	return count
}
