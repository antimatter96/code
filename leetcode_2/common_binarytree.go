package main

import (
	"fmt"
	"math"
)

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func arrayToBSTTreeNode(arr []int) *TreeNode {
	root := &TreeNode{Val: arr[0]}

	for i := 1; i < len(arr); i++ {
		insertIntoBST(root, &TreeNode{Val: arr[i]})
	}

	return root

}

func printBST(root *TreeNode, s string) {
	if root == nil {
		return
	}
	printBST(root.Left, s+s)
	fmt.Println(s, root.Val)
	printBST(root.Right, s+s)

}

func insertIntoBST(root *TreeNode, insert *TreeNode) {
	if root.Val > insert.Val {
		if root.Left != nil {
			insertIntoBST(root.Left, insert)
		} else {
			root.Left = insert
		}
	} else if (*root).Val < insert.Val {
		if root.Right != nil {
			insertIntoBST(root.Right, insert)
		} else {
			root.Right = insert
		}
	}
}

func findInBST(root *TreeNode, target int) *TreeNode {
	r := root

	for root != nil && r.Val != target {
		if r.Val > target {
			r = r.Left
		} else if r.Val < target {
			r = r.Right
		}
	}

	return r
}

func insertIntoBinary(arr []int, index int) *TreeNode {

	if index >= len(arr) {
		return nil
	}

	if arr[index] == math.MinInt32 {
		return nil
	}

	treeNode := &TreeNode{Val: arr[index]}

	treeNode.Left = insertIntoBinary(arr, (2*index)+1)
	treeNode.Right = insertIntoBinary(arr, (2*index)+2)

	return treeNode
}

func arrayToBinaryTreeNode(arr []int) *TreeNode {
	return insertIntoBinary(arr, 0)
}

func binaryTreeNodeToArray(root *TreeNode) []int {
	mp := make(map[int]int)

	_binaryTreeNodeToArray(root, mp, 0)

	max := 0

	for k := range mp {
		if k > max {
			max = k
		}
	}

	arr := make([]int, max+1)

	for k, v := range mp {
		arr[k] = v
	}

	return arr
}

func _binaryTreeNodeToArray(root *TreeNode, mp map[int]int, index int) {
	if root == nil {
		return
	}

	mp[index] = root.Val

	if root.Left != nil {
		_binaryTreeNodeToArray(root.Left, mp, (index*2)+1)
	}

	if root.Right != nil {
		_binaryTreeNodeToArray(root.Right, mp, (index*2)+2)
	}
}
