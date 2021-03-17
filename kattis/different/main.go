package main

import (
	"fmt"
)

func main() {
	var x, y int64

	for {
		_, err := fmt.Scanf("%d %d", &x, &y)
		if err != nil {
			break
		}

		if x > y {
			fmt.Println(x - y)
		} else {
			fmt.Println(y - x)
		}

	}
}
