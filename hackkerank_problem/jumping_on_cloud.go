package hackkerankproblem

func JumpingOnClouds(c []int32, k int32) int32 {
	var energy int32 = 100
	var i int32 = 0
	for {
		i = (i + k) % int32(len(c))
		energy--
		if c[i] == 1 {
			energy -= 2
		}
		if i == 0 {
			break
		}
	}
	return energy
}
