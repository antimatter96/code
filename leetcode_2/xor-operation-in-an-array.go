package main

func xorOperation(n int, start int) int {
	result := start

	for i := 1; i < n; i++ {
		start += 2
		result ^= start
	}

	return result
}

func driver__xorOperation() {

}
