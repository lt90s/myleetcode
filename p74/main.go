package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	fmt.Println(searchMatrix(matrix, 3))
	fmt.Println(searchMatrix(matrix, 15))
	fmt.Println(searchMatrix(matrix, 50))
}
