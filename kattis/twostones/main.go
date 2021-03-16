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

	if n%2 == 0 {
		fmt.Println("Bob")
	} else {
		fmt.Println("Alice")
	}

}
