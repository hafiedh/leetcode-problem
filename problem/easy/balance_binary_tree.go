package problem

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(Height(root.Left)-Height(root.Right)) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(Height(root.Left), Height(root.Right)) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
