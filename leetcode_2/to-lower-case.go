package main

import "fmt"

func toLowerCase(str string) string {
	ans := make([]byte, len(str))
	copy(ans, []byte(str))

	for i := 0; i < len(str); i++ {
		if ans[i] > 65-1 && ans[i] < 90+1 {
			ans[i] += 32
		}
	}

	return string(ans)
}

func driver__toLowerCase() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
	fmt.Println(toLowerCase("al&phaBET"))
}
