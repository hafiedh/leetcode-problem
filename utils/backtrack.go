package utils

func Backtrack(result *[][]int, temp []int, candidates []int, remain int, start int) {
	if remain < 0 {
		return
	} else if remain == 0 {
		*result = append(*result, append([]int{}, temp...))
	} else {
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			Backtrack(result, temp, candidates, remain-candidates[i], i)
			temp = temp[:len(temp)-1] // remove last element
		}
	}
}
