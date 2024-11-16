package medium

func IsValidSudoku(board [][]byte) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if !isValidRow(board, i) {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if !isValidColumn(board, i) {
			return false
		}
	}

	// Check 3x3
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			if !isValid3x3(board, i, j) {
				return false
			}
		}
	}

	return true
}

func isValidRow(board [][]byte, row int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[row][i] == '.' {
			continue
		}
		if m[board[row][i]] {
			return false
		}
		m[board[row][i]] = true
	}
	return true
}

func isValidColumn(board [][]byte, column int) bool {
	m := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		if board[i][column] == '.' {
			continue
		}
		if m[board[i][column]] {
			return false
		}
		m[board[i][column]] = true
	}
	return true
}

func isValid3x3(board [][]byte, row, column int) bool {
	m := make(map[byte]bool)
	for i := row; i < row+3; i++ {
		for j := column; j < column+3; j++ {
			if board[i][j] == '.' {
				continue
			}
			if m[board[i][j]] {
				return false
			}
			m[board[i][j]] = true
		}
	}
	return true
}
