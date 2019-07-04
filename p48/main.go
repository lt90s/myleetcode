package main

import "fmt"

func rotate(matrix [][]int) {
	var n int
	if n = len(matrix); n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		for j := i; j < n-i-1; j++ {
			ii := n - 1 - i
			jj := n - 1 - j

			matrix[i][j], matrix[j][ii], matrix[ii][jj], matrix[jj][i] =
				matrix[jj][i], matrix[i][j], matrix[j][ii], matrix[ii][jj]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
