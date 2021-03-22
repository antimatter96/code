package main

func largestAltitude(gain []int) int {
	current := 0

	max := 0

	for i := 0; i < len(gain); i++ {
		current += gain[i]

		if current > max {
			max = current
		}
	}

	return max
}
