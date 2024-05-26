package problem

func InorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]    // pop
		stack = stack[:len(stack)-1] // pop
		res = append(res, cur.Val)   // visit
		cur = cur.Right              // push
	}
	return res
}
