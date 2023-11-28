package problem

func MissingNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
		result ^= i
	}
	result ^= len(nums)
	return result
}
