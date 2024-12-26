package easy

func TargetIndices(nums []int, target int) []int {
	sortedNums := bubbleSort(nums)
	result := []int{}
	for i, num := range sortedNums {
		if num == target {
			result = append(result, i)
		}
	}

	return result
}

func bubbleSort(elements []int) []int {
	keepWorking := true
	for keepWorking {
		keepWorking = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepWorking = true
			}
		}
	}
	return elements
}
