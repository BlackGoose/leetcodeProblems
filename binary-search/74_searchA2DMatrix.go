package main

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	left, right := 0, n*m-1

	for left <= right {
		middle := left + (right-left)/2
		x := middle / m
		y := middle % m
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return false
}
