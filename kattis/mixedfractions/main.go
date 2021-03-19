package main

import (
	"fmt"
)

func main() {
	var a, b int

	for {
		_, err := fmt.Scanf("%d %d", &a, &b)
		if err != nil {
			break
		}

		if a == 0 && b == 0 {
			break
		}

		fmt.Println(int(a/b), a%b, "/", b)
	}
}
