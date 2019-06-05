package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			for k := 0; k < len(matrix) - i && k < len(matrix[i]) - j; k++ {
				ok := true
				for x := 0; x <= k; x++ {
					if matrix[i+k][j+x] == '0' {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
				for x := 0; x <= k; x++ {
					if matrix[i+x][j+k] == '0' {
						ok = false
						break
					}
				}
				if !ok {
					break
				}

				if k + 1 > max {
					max = k + 1
				}
			}
		}
	}
	return max * max
}

func main() {
	fmt.Println(maximalSquare([][]byte{
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}
