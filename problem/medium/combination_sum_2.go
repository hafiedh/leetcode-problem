package medium

import (
	"sort"

	"hafiedh.com/leetcode/utils"
)

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationSum2Recursive(candidates, target)
}

func combinationSum2Recursive(candidates []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		candidate := candidates[i]
		if target == candidate {
			result = append(result, []int{candidate})
		} else if target > candidate {
			subResult := combinationSum2Recursive(candidates[i+1:], target-candidate)
			for _, sub := range subResult {
				result = append(result, append([]int{candidate}, sub...))
			}
		}
	}
	result = utils.RemoveDuplicate2D(result)

	return result
}
