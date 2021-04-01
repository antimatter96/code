package main

import "fmt"

func minOperations2(box string) []int {
	n := len(box)
	boxes := make([]bool, n)

	for i, c := range box {
		if c == 49 {
			boxes[i] = true
		}
	}

	var totalTillNow int

	left := make([]int, n)
	totalTillNow = 0
	if boxes[0] {
		totalTillNow++
	}
	for i := 1; i < n; i++ {
		left[i] = totalTillNow + left[i-1]
		if boxes[i] {
			totalTillNow++
		}
	}

	right := make([]int, len(box))
	totalTillNow = 0
	if boxes[n-1] {
		totalTillNow++
	}
	for i := n - 2; i > -1; i-- {
		right[i] = totalTillNow + right[i+1]
		if boxes[i] {
			totalTillNow++
		}
	}

	total := make([]int, len(box))
	for i := 0; i < len(box); i++ {
		total[i] = left[i] + right[i]
	}

	return total
}

func driver__minOperations() {
	//[1,0,1,0,0,1]
	//[0,1,2,4,6,8]

	//[1,0,1,0,1,1]
	//[0,1,2,4,6,9]

	//[0,1,3]
	//[1,0,0]

	//[0 ,0,0,1,2,4]
	//[11,8,5,3,1,0]

	fmt.Println(minOperations("110"))
	fmt.Println(minOperations("001011"))

}

func minOperations(box string) []int {
	n := len(box)
	boxes := make([]bool, n)

	var totalTillNow int
	left := make([]int, n)
	for i, c := range box {
		if i > 0 {
			left[i] = totalTillNow + left[i-1]
		}

		if c == 49 {
			boxes[i] = true
			totalTillNow++
		}
	}

	right := make([]int, len(box))
	totalTillNow = 0
	if boxes[n-1] {
		totalTillNow++
	}
	total := make([]int, len(box))
	for i := n - 2; i > -1; i-- {
		right[i] = totalTillNow + right[i+1]
		if boxes[i] {
			totalTillNow++
		}
		total[i] = left[i] + right[i]
	}
	total[n-1] = left[n-1]

	return total
}
