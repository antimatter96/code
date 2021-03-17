package main

import (
	"fmt"
)

func main() {
	quantity := []int{1, 1, 2, 2, 2, 8}
	var N int = 6

	var temp int
	for i := 0; i < N; i++ {
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		fmt.Print(quantity[i]-temp, " ")
	}

}
