package medium

import (
	"hafiedh.com/leetcode/utils"
)

func PermuteUnique(nums []int) [][]int {
	result := [][]int{}
	sorted := utils.BubbleSort(nums)
	backtrackPermuteUnique(&result, []int{}, sorted, make([]bool, len(nums)))
	return result
}

func backtrackPermuteUnique(result *[][]int, temp, nums []int, used []bool) {
	if len(temp) == len(nums) {
		*result = append(*result, append([]int{}, temp...))
	} else {
		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
				continue
			}
			used[i] = true
			temp = append(temp, nums[i])
			backtrackPermuteUnique(result, temp, nums, used)
			temp = temp[:len(temp)-1] //
			used[i] = false
		}
	}
}
