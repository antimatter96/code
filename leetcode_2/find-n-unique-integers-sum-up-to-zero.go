package main

func sumZero(n int) []int {
	arr := make([]int, n)

	if n == 1 {
		arr[0] = 0
		return arr
	}

	for i := 0; i < n-1; i++ {
		arr[i] = i + 1
	}

	arr[n-1] = -(n - 1) * (n) / 2

	return arr
}
