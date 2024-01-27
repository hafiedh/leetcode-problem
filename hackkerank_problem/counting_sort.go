package hackkerankproblem

import "fmt"

func CountingSort(arr []int32) []int32 {
	var result []int32
	var countArr [100]int32

	for _, val := range arr {
		fmt.Println("Value", val)
		countArr[val]++
	}

	for _, val := range countArr {
		result = append(result, val)
	}

	return result
}
