package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, e := range candies {
		if max < e {
			max = e
		}
	}
	arr := make([]bool, len(candies))
	for i, e := range candies {
		if e+extraCandies >= max {
			arr[i] = true
		}
	}
	return arr

}

func driver__kidsWithCandies() {

}
