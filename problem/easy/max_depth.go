package problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
