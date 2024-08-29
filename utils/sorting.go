package utils

func BubbleSort(elements []int) []int {
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

func QuickSort(elements []int) []int {
	if len(elements) <= 1 {
		return elements
	}

	n := len(elements) / 2
	pivot := elements[n]

	left := make([]int, 0)
	right := make([]int, 0)

	for i := range elements {
		if i == n {
			continue
		}
		if elements[i] <= pivot {
			left = append(left, elements[i])
		} else {
			right = append(right, elements[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	left = append(left, pivot)
	left = append(left, right...)

	return left
}
