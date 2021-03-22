package main

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	res := 0

	for i := 0; i < n; i++ {
		temp := 0
		for j := i; j < n; j++ {
			temp += arr[j]
			l := (j - i + 1)
			if l%2 == 1 {
				res += temp
			}
		}

	}
	return res
}

func sumOddLengthSubarrays2(arr []int) int {
	n := len(arr)
	res := 0
	for i := 0; i < n; i++ {
		// k = (i+1)*(n-i)
		// (k + 1)/2
		// k/2
		res += ((i+1)*(n-i) + 1) / 2 * arr[i]
	}
	return res
}
