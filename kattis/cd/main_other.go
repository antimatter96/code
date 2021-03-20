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

		mp := make(map[int]bool, n)

		for i := 0; i < n; i++ {
			scanner.Scan()
			temp = toInt(scanner.Bytes())
			mp[temp] = true
		}

		for i := 0; i < m; i++ {
			scanner.Scan()
			temp = toInt(scanner.Bytes())

			delete(mp, temp)
		}

		fmt.Println(n - len(mp))
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
