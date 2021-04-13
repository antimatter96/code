package main

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	mp := make(map[int]map[int]bool)
	// 	mp := map[int]map[int]bool{} is faster

	for _, log := range logs {
		userId := log[0]
		minute := log[1]

		if mp[userId] == nil {
			mp[userId] = make(map[int]bool)
			// mp[userId] = map[int]bool{} is faster
		}

		mp[userId][minute] = true

	}
	arr := make([]int, k)

	for _, minutes := range mp {
		arr[len(minutes)-1]++
	}

	return arr
}

func driver__findingUsersActiveMinutes() {
	test_cases := []struct {
		logs   [][]int
		k      int
		output []int
	}{
		{[][]int{
			{0, 5}, {1, 2},
			{0, 2}, {0, 5},
			{1, 3},
		}, 5, []int{0, 2, 0, 0, 0}},
		{[][]int{
			{1, 1}, {2, 2}, {2, 3},
		}, 4, []int{1, 1, 0, 0}},
	}

	for _, tc := range test_cases {
		fmt.Println(findingUsersActiveMinutes(tc.logs, tc.k), tc.output)
	}
}
