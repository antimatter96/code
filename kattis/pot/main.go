package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	var sum int64
	var num int64
	var pow int64
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		pow = (num % 10)
		num = num / 10

		sum += int64(math.Pow(float64(num), float64(pow)))
	}

	fmt.Printf("%d\n", sum)

}
