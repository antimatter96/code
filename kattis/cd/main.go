package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var n, m int

		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d %d", &n, &m)

		if n == 0 && m == 0 {
			break
		}

		var temp int

		count := 0
		arr := make([]int, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			arr[i] = toInt(scanner.Bytes())
		}

		cursor := 0
		for i := 0; i < m; i++ {
			scanner.Scan()
			temp = toInt(scanner.Bytes())

			for cursor < n {

				if arr[cursor] > temp {
					break
				} else if arr[cursor] < temp {
					cursor++
				} else {
					count++
					cursor++
					break
				}
			}

		}

		fmt.Println(count)
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
