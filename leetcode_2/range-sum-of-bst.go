package main

// func rangeSumBST(root *TreeNode, low int, high int) int {
// 	var ans int

// 	rec(root, low, high, &ans)

// 	return ans
// }

// func rec(root *TreeNode, low int, high int, sum *int) {
// 	if root == nil {
// 		return
// 	}

// 	if low <= root.Val {
// 		rec(root.Left, low, high, sum)
// 	}

// 	if root.Val >= low && root.Val <= high {
// 		*sum += root.Val
// 	}

// 	if high >= root.Val {
// 		rec(root.Right, low, high, sum)
// 	}
// }

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var ans int
	if low <= root.Val {
		ans += rangeSumBST(root.Left, low, high)
	}

	if root.Val >= low && root.Val <= high {
		ans += root.Val
	}
	if high >= root.Val {
		ans += rangeSumBST(root.Right, low, high)
	}

	return ans
}
