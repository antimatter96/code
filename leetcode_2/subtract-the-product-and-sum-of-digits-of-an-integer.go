package main

import "fmt"

func subtractProductAndSum(n int) int {
	pro := 1
	sum := 0

	var temp int
	for n != 0 {

		temp = n % 10

		sum += temp
		pro *= temp

		n /= 10
	}

	return pro - sum
}

func driver__subtractProductAndSum() {
	fmt.Println(subtractProductAndSum(234) == 15)
	fmt.Println(subtractProductAndSum(4421) == 21)
}
