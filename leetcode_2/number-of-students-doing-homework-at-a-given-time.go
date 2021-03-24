package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	count := 0

	n := len(startTime)

	for i := 0; i < n; i++ {

		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			count++
		}
	}

	return count
}
