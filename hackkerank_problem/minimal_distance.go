package hackkerankproblem

func MinimalDistance(a []int32) int32 {
	var result int32 = -1
	var m map[int32]int32 = make(map[int32]int32)
	for i := 0; i < len(a); i++ {
		if _, ok := m[a[i]]; ok {
			if result == -1 {
				result = int32(i - int(m[a[i]]))
			} else {
				result = checkMinimalDistance(result, int32(i-int(m[a[i]])))
			}
		} else {
			m[a[i]] = int32(i)
		}
	}

	return result
}

func checkMinimalDistance(a, b int32) int32 {
	if a > b {
		return b
	}

	return a
}
