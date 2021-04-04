package main

import "fmt"

func countGoodRectangles(rectangles [][]int) int {
	max := 0
	count := 0
	localMax := 0

	for i := 0; i < len(rectangles); i++ {
		localMax = rectangles[i][0]
		if localMax > rectangles[i][1] {
			localMax = rectangles[i][1]
		}

		if localMax == max {
			count++
		} else if localMax > max {
			max = localMax
			count = 1
		}
	}

	return count
}

func driver__countGoodRectangles() {
	fmt.Println(countGoodRectangles([][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}))
	fmt.Println(countGoodRectangles([][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}))
}
