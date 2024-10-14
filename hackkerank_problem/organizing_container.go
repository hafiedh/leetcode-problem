package hackkerankproblem

func OrganizingContainers(container [][]int32) string {
	// Write your code here
	n := len(container)
	balls := make([]int32, n)
	containers := make([]int32, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			balls[i] += container[i][j]
			containers[j] += container[i][j]
		}
	}
	for i := 0; i < n; i++ {
		found := false
		for j := 0; j < n; j++ {
			if balls[i] == containers[j] {
				containers[j] = -1
				found = true
				break
			}
		}
		if !found {
			return "Impossible"
		}
	}
	return "Possible"
}

// input := [][]int32{{0, 2, 1}, {1, 1, 1}, {2, 0, 0}}
// balls := [3]int32{3, 3, 2}
// containers := [3]int32{3, 3, 2}
