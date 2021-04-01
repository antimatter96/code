package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}

	ptr := head

	for i := 1; i < len(arr); i++ {
		ptr.Next = &ListNode{Val: arr[i]}
		ptr = ptr.Next
	}

	return head
}

func printLinkedList(head *ListNode) {
	ptr := head

	for ptr != nil {
		fmt.Print(ptr.Val, " -> ")
		ptr = ptr.Next
	}
	fmt.Println()
}
