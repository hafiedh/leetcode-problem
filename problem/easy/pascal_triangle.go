package easy

func PascalTriangle(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	var result [][]int
	for i := 0; i < numRows; i++ {
		var row []int
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else {
				row = append(row, result[i-1][j-1]+result[i-1][j]) // add previous row's elements together
			}
		}
		result = append(result, row)
	}
	return result
}
