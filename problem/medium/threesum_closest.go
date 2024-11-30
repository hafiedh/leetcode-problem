package medium

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
