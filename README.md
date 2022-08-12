# golang_valid_sudoku

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.

**Note:**

- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

## Examples

**Example 1:**

![https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png](https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/Sudoku-by-L2G-20050714.svg/250px-Sudoku-by-L2G-20050714.svg.png)

```
Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true

```

**Example 2:**

```
Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the5 in the top left corner being modified to8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

```

**Constraints:**

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit `1-9` or `'.'`.

## 解析

給定一個 9 by 9 的byte 矩陣 board 

每個 board[i][j] = ‘1’ … ‘9’ or ‘.’

如果是 ‘.’ 代表沒放值

定義一個 board is valid 需要符合以下規則

1. 每個 row 不出現重複的 byte 假設 byte ！= ‘.’
2. 每個 column 不出現重複的 byte 假設 byte ！= ‘.’
3. 把 board分割成 9個  3 x 3 小方格 mask, mask 內不出現重複的 byte 假設 byte ！= ‘.’

要求寫一個演算法來判斷這個 board 是不是 valid

透過 建立3個 hashTable rows, columns, masks 分別用來紀錄已經出現過的值

rows 用來紀錄某個row已經出現過的 0-9

columns 用來紀錄某個columns已經出現過的 0-9

masks 用來紀錄某個mask已經出現過的 0-9

值得一提的是, 可以透過整數除法的方式把每個 row, col 對應到 不同的 mask 如下

![](https://i.imgur.com/r1pzW3F.png)

當發現某個 hashTable 出現重複的值則回傳 false

當所有值都檢查完且沒有重複

則回傳 true

因為固定是 9 by 9 board 所以時間複雜度跟空間複雜度都是固定值

時間複雜度是 O(1)

空間複雜度是 O(1)

## 程式碼
```go
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

```
## 困難點

1. 需要想出如何把 row, col 對應到 9 個切分小方格的方式

## Solve Point

- [x]  建立3個 hashTable, rows, cols, masks
- [x]  從 row := 0..8, col := 0..8 逐步做以下檢查
- [x]  如果 rows[row][board[row][col]] 存在 則回傳 false, 否則把 rows[row][board[row][col]] = struct{}{}
- [x]  如果 cols[col][board[row][col]] 存在 則回傳 false, 否則把 cols[col][board[row][col]] = struct{}{}
- [x]  如果 masks[row/3, col/3][board[row][col]] 存在 則回傳 false, 否則把 masks[row/3, col/3][board[row][col]] = struct{}{}
- [x]  當所有 row, col 都走完且都沒有重複 則回傳 true