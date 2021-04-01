package main

import "fmt"

func removeDuplicates(S string) string {
	arr := make([]byte, 0, len(S))

	var n int

	for i := 0; i < len(S); i++ {
		n = len(arr)

		if n == 0 {
			arr = append(arr, S[i])
		} else if arr[n-1] == S[i] {
			arr = arr[:n-1]
		} else {
			arr = append(arr, S[i])
		}
	}

	return string(arr)
}

func driver__removeDuplicates() {
	fmt.Println(removeDuplicates("abbaca"))
}
