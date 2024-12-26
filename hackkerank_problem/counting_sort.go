package hackkerankeasy

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

func CheckInMap(arr []int32) int32 {
	result := make(map[int32]int32)
	for _, val := range arr {
		result[val]++
	}

	var max int32
	for _, val := range result {
		if val > max {
			max = val
		}
	}

	return max

}
