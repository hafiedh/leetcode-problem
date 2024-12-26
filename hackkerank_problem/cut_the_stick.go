package hackkerankeasy

func CutTheSticks(arr []int32) []int32 {
	var result []int32
	for len(arr) > 0 {
		result = append(result, int32(len(arr)))
		min := arr[0]
		for _, v := range arr {
			if v < min {
				min = v
			}
		}
		var temp []int32
		for _, v := range arr {
			if v != min {
				temp = append(temp, v-min)
			}
		}
		arr = temp
	}
	return result
}
