package main

import (
	"fmt"
)

func main() {
	var inp string

	_, err := fmt.Scanf("%s", &inp)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("%c", inp[0])
	for i := 1; i < len(inp)-1; i++ {
		if inp[i] == '-' {
			fmt.Printf("%c", inp[i+1])
		}
	}

}
