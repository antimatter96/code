package main

import (
	"fmt"
)

func main() {
	var n int

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var temp int
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		if temp%2 == 0 {
			fmt.Println(temp, "is even")
		} else {
			fmt.Println(temp, "is odd")
		}
	}

}
