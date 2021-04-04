package main

import (
	"fmt"
	"strings"
)

// 8 ms
func sortString(s string) string {
	var order string = "abcdefghijklmnopqrstuvwxyz"
	var reversed string = "zyxwvutsrqponmlkjihgfedcba"
	mp := make(map[rune]int, 26)

	for _, r := range s {
		mp[r]++
	}

	count := len(s)
	ans := &strings.Builder{}
	ans.Grow(count)

	for count > 0 {
		for _, r := range order {
			if mp[r] > 0 {
				mp[r]--

				ans.WriteRune(r)
				count--
			}
		}

		for _, r := range reversed {
			if mp[r] > 0 {
				mp[r]--

				ans.WriteRune(r)
				count--
			}
		}

	}

	return ans.String()
}

// 0ms
func sortString2(s string) string {
	var order string = "abcdefghijklmnopqrstuvwxyz"
	var reversed string = "zyxwvutsrqponmlkjihgfedcba"
	mp := make([]int, 26)

	for _, r := range s {
		mp[r-97]++
	}

	count := len(s)
	ans := &strings.Builder{}
	ans.Grow(count)

	for count > 0 {
		for _, r := range order {
			if mp[r-97] > 0 {
				mp[r-97]--

				ans.WriteRune(r)
				count--
			}
		}

		for _, r := range reversed {
			if mp[r-97] > 0 {
				mp[r-97]--

				ans.WriteRune(r)
				count--
			}
		}

	}

	return ans.String()
}

func driver__sortString() {
	fmt.Println(sortString2("aaaabbbbcccc") == "abccbaabccba")
	fmt.Println(sortString2("rat") == "art")
	fmt.Println(sortString2("leetcode") == "cdelotee")
	fmt.Println(sortString2("ggggggg") == "ggggggg")
	fmt.Println(sortString2("spo") == "ops")
}
