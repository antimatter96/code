package main

import "fmt"

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

func driver__countGoodTriplets() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
	fmt.Println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))
	fmt.Println(countGoodTriplets([]int{7, 3, 7, 3, 12, 1, 12, 2, 3}, 5, 8, 1))
}
