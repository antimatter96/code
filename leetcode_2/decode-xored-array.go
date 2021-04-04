package main

import "fmt"

func decode(encoded []int, first int) []int {
	arr := make([]int, len(encoded)+1)

	arr[0] = first
	for i := 1; i < len(encoded)+1; i++ {
		arr[i] = encoded[i-1] ^ arr[i-1]
	}

	return arr
}

func driver__decode() {
	test_cases := []struct {
		encoded []int
		first   int
		answer  []int
	}{
		{encoded: []int{1, 2, 3}, first: 1, answer: []int{1, 0, 2, 1}},
		{encoded: []int{6, 2, 7, 3}, first: 4, answer: []int{4, 2, 0, 7, 4}},
	}

	for _, tc := range test_cases {
		fmt.Println(decode(tc.encoded, tc.first), tc.answer)
	}

}
