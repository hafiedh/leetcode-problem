package utils

func RemoveDuplicate2D(elements [][]int) [][]int {
	encountered := map[string]bool{}
	result := [][]int{}
	for _, element := range elements {
		key := ""
		for _, e := range element {
			key += string(e)
		}
		if !encountered[key] {
			encountered[key] = true
			result = append(result, element)
		}
	}
	return result
}

func RemoveDuplicate(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}
	for _, element := range elements {
		if !encountered[element] {
			encountered[element] = true
			result = append(result, element)
		}
	}
	return result
}
