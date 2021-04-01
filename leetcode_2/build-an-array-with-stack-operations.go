package main

import "fmt"

func buildArray(target []int, n int) []string {
	const push string = "Push"
	const pop string = "Pop"
	res := make([]string, 0, 2*len(target))

	toAdd := 0

	for i := 1; i < n+1 && toAdd < len(target); i++ {
		if i == target[toAdd] {
			toAdd++
			res = append(res, push)
		} else if i < target[toAdd] {
			res = append(res, push, pop)
		} else {
			toAdd++
		}
	}

	return res
}

func driver__buildArray() {
	fmt.Println(buildArray([]int{1, 3}, 3))
	fmt.Println(buildArray([]int{1, 2, 3}, 3))
	fmt.Println(buildArray([]int{1, 2}, 4))
	fmt.Println(buildArray([]int{2, 3, 4}, 4))
}
