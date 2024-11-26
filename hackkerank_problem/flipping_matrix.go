package hackkerankeasy

func FlippingMatrix(matrix [][]int32) int32 {
	var result int32

	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			result += max(matrix[i][j], matrix[i][n-j-1], matrix[n-i-1][j], matrix[n-i-1][n-j-1])
		}
	}

	return result
}

func max(a, b, c, d int32) int32 {
	if a > b {
		b = a
	}

	if c > d {
		d = c
	}

	if b > d {
		return b
	}

	return d
}
