package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	_, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil {
		panic(err)
	}

	if a+b == c {
		fmt.Printf("%d+%d=%d", a, b, c)
	} else if a == b+c {
		fmt.Printf("%d=%d+%d", a, b, c)
	} else if a-b == c {
		fmt.Printf("%d-%d=%d", a, b, c)
	} else if a == b-c {
		fmt.Printf("%d=%d-%d", a, b, c)
	} else if a*b == c {
		fmt.Printf("%d*%d=%d", a, b, c)
	} else if a == b*c {
		fmt.Printf("%d=%d*%d", a, b, c)
	} else if a/b == c {
		fmt.Printf("%d/%d=%d", a, b, c)
	} else if a == b/c {
		fmt.Printf("%d=%d/%d", a, b, c)
	}
}
