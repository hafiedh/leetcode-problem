package hackkerankeasy

func NonDivisibleSubset(k int32, s []int32) int32 {
	var count int32 = 0
	var remainder = make([]int32, k)
	for _, v := range s {
		remainder[v%k]++
	}
	if remainder[0] > 0 {
		count++
	}
	for i := int32(1); i <= k/2; i++ {
		if i != k-i {
			count += maxSubset(remainder[i], remainder[k-i])
		} else {
			count++
		}
	}
	return count
}

func maxSubset(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
