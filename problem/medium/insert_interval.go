package medium

func Insert(intervals [][]int, newInterval []int) [][]int {
	// Initialize an empty result slice
	result := [][]int{}

	// Initialize the current interval to the new interval
	currentInterval := newInterval

	// Iterate through the intervals
	for _, interval := range intervals {
		// If the current interval does not overlap with the new interval, add it to the result
		if interval[1] < currentInterval[0] {
			result = append(result, interval)
		} else if interval[0] > currentInterval[1] {
			// If the current interval is completely before the new interval, add the new interval to the result
			result = append(result, currentInterval)
			currentInterval = interval
		} else {
			// If there is an overlap, merge the intervals
			currentInterval[0] = min(currentInterval[0], interval[0])
			currentInterval[1] = max(currentInterval[1], interval[1])
		}
	}

	// Add the last merged interval to the result
	result = append(result, currentInterval)

	return result
}
