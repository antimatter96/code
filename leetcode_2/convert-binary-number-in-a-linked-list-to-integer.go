package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	count := 0
	temp := head

	for temp != nil {
		count *= 2
		count += temp.Val
		temp = temp.Next
	}

	return count
}
