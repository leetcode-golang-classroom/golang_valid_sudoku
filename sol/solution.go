package sol

import "fmt"

func isValidSudoku(board [][]byte) bool {
	ROWS, COLS := len(board), len(board[0])
	rows, cols := make(map[int](map[byte]struct{})), make(map[int](map[byte]struct{}))
	masks := make(map[string](map[byte]struct{}))
	var checkValid = func(row, col int) bool {
		masksRow, masksCol := row/3, col/3
		masksKey := fmt.Sprintf("%d%d", masksRow, masksCol)
		val := board[row][col]
		if _, ok := masks[masksKey][val]; ok {
			return false
		}
		if _, ok := rows[row][val]; ok {
			return false
		}
		if _, ok := cols[col][val]; ok {
			return false
		}
		if _, ok := masks[masksKey]; !ok {
			masks[masksKey] = make(map[byte]struct{})
		}
		masks[masksKey][val] = struct{}{}
		if _, ok := rows[row]; !ok {
			rows[row] = make(map[byte]struct{})
		}
		rows[row][val] = struct{}{}
		if _, ok := cols[col]; !ok {
			cols[col] = make(map[byte]struct{})
		}
		cols[col][val] = struct{}{}
		return true
	}
	for row := 0; row < ROWS; row++ {
		for col := 0; col < COLS; col++ {
			if board[row][col] == '.' {
				continue
			}
			if checkValid(row, col) == false {
				return false
			}
		}
	}
	return true
}
