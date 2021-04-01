package main

func numJewelsInStones(jewels string, stones string) int {
	mp := make(map[rune]bool)

	for _, c := range jewels {
		mp[c] = true
	}

	count := 0
	for _, c := range stones {
		if mp[c] {
			count++
		}
	}

	return count
}

func driver__numJewelsInStones() {

}
