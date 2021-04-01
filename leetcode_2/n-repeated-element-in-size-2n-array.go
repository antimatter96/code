package main

func repeatedNTimes(A []int) int {
	mp := make(map[int]bool)

	for _, v := range A {
		if mp[v] {
			return v
		} else {
			mp[v] = true
		}
	}

	return -1
}

func driver__repeatedNTimes() {

}
