package hackkerankproblem

func ChocolateFeast(n int32, c int32, m int32) int32 {
	// Write your code here
	var result int32 = n / c
	var wrapper int32 = result
	for wrapper >= m {
		result += wrapper / m
		wrapper = wrapper/m + wrapper%m
	}
	return result
}
