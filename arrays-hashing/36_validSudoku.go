package main

type square struct {
	x, y int
}

func isValidSudoku(board [][]byte) bool {
	verticalMap := make(map[int][9]bool, 9)
	horizontalMap := make(map[int][9]bool, 9)
	squareMap := make(map[square][9]bool, 9)

	for rowIndex, rows := range board {
		for columnIndex, value := range rows {
			if value == '.' {
				continue
			}
			value -= '0' - 1
			if verticalMap[rowIndex][value] {
				return false
			} else {
				temp := verticalMap[rowIndex]
				temp[value] = true
				verticalMap[rowIndex] = temp
			}
			if horizontalMap[columnIndex][value] {
				return false
			} else {
				temp := horizontalMap[columnIndex]
				temp[value] = true
				horizontalMap[columnIndex] = temp
			}
			if squareMap[square{x: columnIndex / 3, y: rowIndex / 3}][value] {
				return false
			} else {
				temp := squareMap[square{x: columnIndex / 3, y: rowIndex / 3}]
				temp[value] = true
				squareMap[square{x: columnIndex / 3, y: rowIndex / 3}] = temp
			}
		}
	}
	return true
}
