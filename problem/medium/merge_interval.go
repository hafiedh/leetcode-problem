package medium

func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	intervals = sorting(intervals)

	result := [][]int{}
	currentInterval := intervals[0]
	for _, interval := range intervals[1:] {
		if interval[0] > currentInterval[1] {
			result = append(result, currentInterval)
			currentInterval = interval
		} else {
			currentInterval[0] = min(currentInterval[0], interval[0])
			currentInterval[1] = max(currentInterval[1], interval[1])
		}
	}

	result = append(result, currentInterval)
	return result
}

func sorting(interval [][]int) [][]int {
	for i := range interval {
		for j := i + 1; j < len(interval); j++ {
			if interval[i][0] > interval[j][0] {
				interval[i], interval[j] = interval[j], interval[i]
			}
		}
	}
	return interval
}
