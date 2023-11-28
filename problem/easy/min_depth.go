package problem

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}

	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}

	return Min(MinDepth(root.Left), MinDepth(root.Right)) + 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
