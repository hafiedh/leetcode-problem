package medium

import (
	"sort"
)

func FourSum(nums []int, target int) [][]int {
	// Input: nums = [1,0,-1,0,-2,2], target = 0
	// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j != (i+1) && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					temp := []int{nums[i], nums[j], nums[k], nums[l]}
					result = append(result, temp)
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return result
}
