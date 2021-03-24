package main

func findNumbers(nums []int) int {
	noOfDigits := func(x int) int {
		i := 1

		for x > 9 {
			x /= 10
			i++
		}
		return i
	}

	count := 0

	for _, num := range nums {
		if noOfDigits(num)%2 == 0 {
			count++
		}
	}

	return count
}
