package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	_, err := fmt.Scanf("%d\n%d", &x, &y)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if x > 0 && y > 0 {
		fmt.Println(1)
	} else if x > 0 && y < 0 {
		fmt.Println(4)
	} else if x < 0 && y < 0 {
		fmt.Println(3)
	} else if x < 0 && y > 0 {
		fmt.Println(2)
	}

}
