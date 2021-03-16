package main

import (
	"fmt"
)

func main() {
	var avg int
	var r1 int

	_, err := fmt.Scanf("%d %d", &r1, &avg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println((2 * avg) - r1)
}
