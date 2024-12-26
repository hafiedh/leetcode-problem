package hackkerankeasy

func Ways(total, k int32) int32 {
	ways := make([]int, total+1)
	ways[0] = 1
	for i := int32(1); i <= k; i++ {
		for j := i; j <= total-k; j++ {
			ways[j] += ways[j-i]
		}
	}
	return int32(ways[total])
}
