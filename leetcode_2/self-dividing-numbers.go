package main

import "fmt"

func getDigits(n int) []int {
	arr := make([]int, 0)

	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}

	return arr
}

func helper(n int) bool {
	temp := n

	for temp != 0 {
		digit := temp % 10

		if digit == 0 || n%digit != 0 {
			return false
		}

		temp /= 10
	}

	return true
}

func selfDividingNumbers(left int, right int) []int {
	arr := make([]int, 0, right-left+1)

	for i := left; i <= right; i++ {
		if helper(i) {
			arr = append(arr, i)
		}
	}

	return arr
}

func selfDividingNumbers_2(left int, right int) []int {
	arr := make([]int, 0, right-left+1)

	for i := left; i <= right; i++ {
		ints := getDigits(i)

		good := true
		for _, digit := range ints {
			if digit == 0 || i%digit != 0 {
				good = false
				break
			}
		}

		if good {
			arr = append(arr, i)
		}

	}

	return arr
}

func driver__selfDividingNumbers() {
	fmt.Println(selfDividingNumbers(1, 22))
	fmt.Println(selfDividingNumbers(1, 23))
	fmt.Println(selfDividingNumbers(1, 24))
}
