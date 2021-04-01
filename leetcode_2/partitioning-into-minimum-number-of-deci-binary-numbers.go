package main

import "fmt"

func minPartitions(n string) int {
	var k byte

	for i := 0; i < len(n); i++ {
		if k < n[i] {
			k = n[i]
		}

		if k == '9' {
			break
		}
	}

	return int(k) - 48
}

func driver__minPartitions() {
	fmt.Println(minPartitions("32"))
	fmt.Println(minPartitions("82734"))
	fmt.Println(minPartitions("27346209830709182346"))
}
