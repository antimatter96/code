package main

import (
	"fmt"
	"sort"
)

func main() {
	inp := make([]int, 3)

	var order string

	_, err := fmt.Scanf("%d%d%d\n%s", &inp[0], &inp[1], &inp[2], &order)
	if err != nil {
		panic(err)
	}

	sort.Ints(inp)

	for _, c := range order {
		fmt.Print(inp[c-65], " ")
	}
}
