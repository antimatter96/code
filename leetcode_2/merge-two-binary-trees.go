package main

import (
	"fmt"
	"math"
)

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	newNode := &TreeNode{}

	if root1 != nil {
		newNode.Val += root1.Val
	}
	if root2 != nil {
		newNode.Val += root2.Val
	}

	if root1 != nil && root2 != nil {
		newNode.Left = mergeTrees(root1.Left, root2.Left)
		newNode.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil {
		newNode.Left = mergeTrees(root1.Left, nil)
		newNode.Right = mergeTrees(root1.Right, nil)
	} else if root2 != nil {
		newNode.Left = mergeTrees(nil, root2.Left)
		newNode.Right = mergeTrees(nil, root2.Right)
	}

	return newNode
}

func driver__mergeTrees() {
	root1 := arrayToBinaryTreeNode([]int{1, 3, 2, 5})
	root2 := arrayToBinaryTreeNode([]int{2, 1, 3, math.MinInt32, 4, math.MinInt32, 7})

	merged := mergeTrees(root1, root2)

	fmt.Println(binaryTreeNodeToArray(root1))
	fmt.Println(binaryTreeNodeToArray(root2))
	fmt.Println(binaryTreeNodeToArray(merged))

	root1 = arrayToBinaryTreeNode([]int{1})
	root2 = arrayToBinaryTreeNode([]int{1, 2})

	merged = mergeTrees(root1, root2)

	fmt.Println(binaryTreeNodeToArray(root1))
	fmt.Println(binaryTreeNodeToArray(root2))
	fmt.Println(binaryTreeNodeToArray(merged))
}
