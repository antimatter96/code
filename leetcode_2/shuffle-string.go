package main

import "fmt"

func restoreString(s string, indices []int) string {
	n := len(s)
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		ans[indices[i]] = s[i]
	}

	return string(ans)
}

func driver__restoreString() {

	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}) == "leetcode")
	fmt.Println(restoreString("abc", []int{0, 1, 2}) == "abc")
	fmt.Println(restoreString("aiohn", []int{3, 1, 4, 2, 0}) == "nihao")
	fmt.Println(restoreString("aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}) == "arigatou")
	fmt.Println(restoreString("art", []int{1, 0, 2}) == "rat")

}
