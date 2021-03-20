package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mp := map[rune]string{
		'a': "@", 'b': "8",
		'c': "(", 'd': "|)",
		'e': "3", 'f': "#",
		'g': "6", 'h': "[-]",
		'i': "|", 'j': "_|",
		'k': "|<", 'l': "1",
		'm': "[]\\/[]", 'n': "[]\\[]",
		'o': "0", 'p': "|D",
		'q': "(,)", 'r': "|Z",
		's': "$", 't': "']['",
		'u': "|_|", 'v': "\\/",
		'w': "\\/\\/", 'x': "}{",
		'y': "`/", 'z': "2",
	}

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()

	for _, c := range s {
		if c >= 65 && c < (65+26) {
			//fmt.Println(c, mp[c+32])
			fmt.Printf("%s", mp[c+32])
		} else if c >= 97 && c < (97+26) {
			//fmt.Println(c, mp[c])
			fmt.Printf("%s", mp[c])
		} else {
			//fmt.Println(c)
			fmt.Printf("%c", c)
		}

	}
	fmt.Println()

}
