package main

import "fmt"

func main() {
	var n int

	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err)
	}

	c := 0
	correction := 0
	for n != 0 {
		c++
		if n != 1 && n%2 == 1 {
			correction = 1
		}
		n /= 2

	}

	fmt.Println(c + correction)
}
