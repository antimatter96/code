package main

import (
	"fmt"
)

func main() {
	var x, n int

	_, err := fmt.Scanf("%d\n%d", &x, &n)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	tots := x
	var temp int
	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &temp)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		tots += (x - temp)
	}

	fmt.Println(tots)

}
