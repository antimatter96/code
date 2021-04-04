package main

import "fmt"

func maximum69Number(num int) int {
	a := 100_000
	ans := 0

	for num != 0 {
		s := num / a
		if s == 6 {
			ans += a * 9
			ans += (num % a)
			break
		} else {
			ans += s * a
		}

		num = num % a
		a = a / 10
	}

	return ans
}

func driver__maximum69Number() {
	fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9996))
	fmt.Println(maximum69Number(9999))

	fmt.Println(maximum69Number(6999))
	fmt.Println(maximum69Number(6699))
}
