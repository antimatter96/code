package main

import (
	"fmt"
	"strings"
)

func freqAlphabets(s string) string {
	bytes := []byte(s)

	for i := 0; i < len(s); i++ {
		if bytes[i] == '#' {
			bytes[i] = 48 + (bytes[i-1] - 48) + (10 * (bytes[i-2] - 48))
			bytes[i-1] = 0
			bytes[i-2] = 0
		}
	}

	ans := &strings.Builder{}
	for i := 0; i < len(s); i++ {
		if bytes[i] != 0 {
			ans.WriteByte((96 + bytes[i] - 48))
		}
	}

	return ans.String()
}

func driver__freqAlphabets() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
	fmt.Println(freqAlphabets("25#"))
	fmt.Println(freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"))
}
