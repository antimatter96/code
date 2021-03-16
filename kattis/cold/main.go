package main

import (
	"fmt"
)

func main() {
	var n int
	var count int

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

		if temp < 0 {
			count++
		}
	}
	fmt.Println(count)

}
