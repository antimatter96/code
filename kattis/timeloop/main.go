package main

import (
	"fmt"
)

func main() {
	var output = "Abracadabra"
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for i := 0; i < n; i++ {
		fmt.Println(i+1, output)
	}
}
