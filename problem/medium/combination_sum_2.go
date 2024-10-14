package medium

import "hafiedh.com/leetcode/utils"

func CombinationSum2(candidates []int, target int) [][]int {
	sortedSlice := utils.QuickSort(candidates)
	return combinationSum2Recursive(sortedSlice, target)
}

func combinationSum2Recursive(candidates []int, target int) [][]int {
	var result [][]int
	for i, candidate := range candidates {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		if target == candidate {
			result = append(result, []int{candidate})
		} else if target > candidate {
			subResult := combinationSum2Recursive(candidates[i+1:], target-candidate)
			for _, sub := range subResult {
				result = append(result, append([]int{candidate}, sub...))
			}
		}
	}
	// remove duplicate
	result = utils.RemoveDuplicate2D(result)

	return result
}
