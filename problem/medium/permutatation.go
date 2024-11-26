package medium

func Permute(nums []int) [][]int {
	result := [][]int{}
	backtrackPermute(&result, []int{}, nums)
	return result
}

func backtrackPermute(result *[][]int, temp, nums []int) {
	if len(temp) == len(nums) {
		*result = append(*result, append([]int{}, temp...))
	} else {
		for i := 0; i < len(nums); i++ {
			if containsInt(temp, nums[i]) {
				continue
			}
			temp = append(temp, nums[i])
			backtrackPermute(result, temp, nums)
			temp = temp[:len(temp)-1]
		}
	}
}

func containsInt(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
