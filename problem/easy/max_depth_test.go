package problem

import "testing"

func TestMaxDepth(t *testing.T) {
	// Test case 1: Empty tree
	if MaxDepth(nil) != 0 {
		t.Errorf("Expected depth: 0, but got %d", MaxDepth(nil))
	}

	root := &TreeNode{Val: 1}
	if MaxDepth(root) != 1 {
		t.Errorf("Expected depth: 1, but got %d", MaxDepth(root))
	}

	// Test case 3: Tree with two levels
	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	if MaxDepth(root) != 2 {
		t.Errorf("Expected depth: 2, but got %d", MaxDepth(root))
	}

	// Test case 4: Tree with three levels
	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	if MaxDepth(root) != 3 {
		t.Errorf("Expected depth: 3, but got %d", MaxDepth(root))
	}
}
