package main

import (
	"fmt"
)

func main() {
	var s1, s2 string

	_, err := fmt.Scanf("%s\n%s", &s1, &s2)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if len(s1) >= len(s2) {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}

}
