package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	abc := "abcdefghijklmnopqrstuvwxyz"
	var n int

	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inp := scanner.Text()

		mp := make(map[rune]bool)
		for _, c := range abc {
			mp[c] = true
		}

		for _, c := range inp {
			delete(mp, c)
			delete(mp, c+32)
		}

		if len(mp) == 0 {
			fmt.Println("pangram")
			continue
		}

		arr := make([]int, 0, len(mp))
		for c := range mp {
			arr = append(arr, int(c)-0)
		}
		fmt.Print("missing ")
		sort.Ints(arr)
		for _, c := range arr {
			fmt.Print(string(c))
		}
		fmt.Println()

	}

}
