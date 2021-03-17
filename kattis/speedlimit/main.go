package main

import (
	"fmt"
)

func main() {
	var n int

	var speed, time int
	var total, prev int
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		if n == -1 {
			break
		}

		total = 0
		prev = 0

		for i := 0; i < n; i++ {
			_, err := fmt.Scanf("%d %d", &speed, &time)
			if err != nil {
				break
			}

			total += speed * (time - prev)
			prev = time
		}

		fmt.Println(total, "miles")
	}
}
