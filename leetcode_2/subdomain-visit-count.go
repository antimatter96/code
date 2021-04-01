package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	mp := make(map[string]int, len(cpdomains))

	for i := 0; i < len(cpdomains); i++ {
		temp := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(temp[0])
		domain := temp[1]

		split := strings.Split(domain, ".")

		s := split[len(split)-1]
		mp[s] += count
		for i := len(split) - 2; i > -1; i-- {
			s = split[i] + "." + s
			mp[s] += count
		}

	}

	response := make([]string, 0, len(mp))

	for k, v := range mp {
		// response = append(response, fmt.Sprintf("%d %s", v, k)) FASTER
		response = append(response, strconv.Itoa(v)+" "+k) // MEMORY
	}

	return response
}

func driver__subdomainVisits() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
