package medium

func CombinationSum(candidates []int, target int) [][]int {
	// example: candidates = [2,3,6,7], target = 7
	// result: [[2,2,3],[7]]
	var result [][]int
	var temp []int
	backtrack(&result, temp, candidates, target, 0)
	return result
}

func backtrack(result *[][]int, temp []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		*result = append(*result, append([]int{}, temp...))
	} else {
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			backtrack(result, temp, candidates, remain-candidates[i], i)
			temp = temp[:len(temp)-1] // remove last element
		}
	}
}
