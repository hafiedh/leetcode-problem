package easy

func FindDuplicate(nums []int) int {
	visited := make(map[int]bool)
	for _, num := range nums {
		if visited[num] {
			return num
		}
		visited[num] = true
	}
	return -1
}
