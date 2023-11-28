package problem

// Given an integer array nums where the elements are sorted in ascending order, convert it to a
// height-balanced
//  binary search tree.

func SortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBST(nums, 0, len(nums)-1)
	// output [0,-10,5,null,-3,null,9]
}

func sortedArrayToBST(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums, start, mid-1)
	root.Right = sortedArrayToBST(nums, mid+1, end)
	return root
}

func SortedArrayToBST2(nums []int) *TreeNode {
	return sortedArrayToBST2(nums, 0, len(nums)-1)
}

func sortedArrayToBST2(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end + 1) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST2(nums, start, mid-1)
	root.Right = sortedArrayToBST2(nums, mid+1, end)
	return root
}
