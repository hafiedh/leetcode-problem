package hackkerankproblem

func JumpingOnClouds(c []int32) int32 {
	var result int32 = 0
	for i := 0; i < len(c)-1; i++ {
		if i+2 < len(c) && c[i+2] == 0 {
			i++
		}
		result++
	}

	return result
}
