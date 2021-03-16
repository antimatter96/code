package main

import (
	"fmt"
)

func main() {
	var n int
	var p int

	_, err := fmt.Scanf("%d %d", &n, &p)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(p)
}
