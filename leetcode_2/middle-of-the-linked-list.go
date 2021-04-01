package main

import "fmt"

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	return slow
}

func driver__middleNode() {
	fmt.Println("1", middleNode(arrayToList([]int{1})))
	fmt.Println("2", middleNode(arrayToList([]int{1, 2})))
	fmt.Println("2", middleNode(arrayToList([]int{1, 2, 3})))
	fmt.Println("3", middleNode(arrayToList([]int{1, 2, 3, 4})))
	fmt.Println("3", middleNode(arrayToList([]int{1, 2, 3, 4, 5})))
	fmt.Println("4", middleNode(arrayToList([]int{1, 2, 3, 4, 5, 6})))
}
