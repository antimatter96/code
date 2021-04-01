package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, balances := range accounts {
		sum := 0
		for _, balance := range balances {
			sum += balance
		}

		if max < sum {
			max = sum
		}
	}
	return max
}

func driver__maximumWealth() {

}
