package main

import "fmt"

var testCaseBest1 []int = []int{5, 4, 3, 2, 1}
var testCaseWorst1 []int = []int{1, 2, 3, 4, 5}
var testCaseBest1000 []int
var testCaseWorst1000 []int

func noOfArrows(arr []int) int {
	n := len(arr)
	count := 0
	height := 0

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			continue
		} else {
			height = arr[i]
			arr[i] = 0
			count++
		}

		height--
		for j := i + 1; height > 0 && j < n; j++ {
			if arr[j] == height {
				arr[j] = 0
				height--
			}
		}
	}

	return count
}

func noOfArrowsNew(arr []int) int {
	n := len(arr)
	count := 0
	height := 0

	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[arr[i]]++
	}

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			continue
		} else {
			height = arr[i]
			arr[i] = 0
			mp[height]--
			count++
		}

		height--
		for j := i + 1; height > 0 && mp[height] > 0 && j < n; j++ {
			if arr[j] == height {
				arr[j] = 0
				mp[height]--
				height--
			}
		}
	}

	return count
}

func init() {
	n := 1000
	testCaseWorst1000 = make([]int, n)
	testCaseBest1000 = make([]int, n)

	for i := 0; i < n; i++ {
		testCaseWorst1000[i] = i + 1
		testCaseBest1000[i] = n - i
	}
}

func main() {
	//fmt.Println(testCaseWorst1000)
	//fmt.Println(testCaseBest1000)

	arr := make([]int, len(testCaseBest1))
	copy(arr, testCaseBest1)
	fmt.Println(noOfArrows(arr))
	copy(arr, testCaseBest1)
	fmt.Println(noOfArrowsNew(arr))

	arr = make([]int, len(testCaseWorst1))
	copy(arr, testCaseWorst1)
	fmt.Println(noOfArrows(arr))
	copy(arr, testCaseWorst1)
	fmt.Println(noOfArrowsNew(arr))
}
