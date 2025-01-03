package medium

// 34. Find First and Last Position of Element in Sorted Array
// Medium
// Topics
// Companies
// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:

// Input: nums = [5,7,7,8,8,9,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums is a non-decreasing array.
// -109 <= target <= 109

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := 0
	right := len(nums) - 1
	start := -1
	end := -1

	for left <= right {
		mid := left + (right-left)/2 // 6 / 2 = 3
		if nums[mid] == target {
			start = mid
			end = mid
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if start == -1 {
		return []int{-1, -1}
	}

	for i := start; i >= 0; i-- {
		if nums[i] == target {
			start = i
		} else {
			break
		}
	}

	for i := end; i < len(nums); i++ {
		if nums[i] == target {
			end = i
		} else {
			break
		}
	}

	return []int{start, end}
}
