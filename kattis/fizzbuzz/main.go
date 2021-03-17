package main

import (
	"fmt"
)

func main() {
	var x, y, n int

	_, err := fmt.Scanf("%d %d %d", &x, &y, &n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for i := 1; i < n+1; i++ {
		if i%x == 0 {
			if i%y == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%y == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
