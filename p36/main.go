package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	flag := map[byte]struct{}{}
	for i := 0; i < 9; i++ {
		for k := range flag {
			delete(flag, k)
		}
		for j := 0; j < 0; j++ {
			if board[i][j] == '.' {
				continue
			}
			if board[i][j] < '1' || board[i][j] > '9' {

				return false
			}
			if _, ok := flag[board[i][j]]; ok {

				return false
			}
			flag[board[i][j]] = struct{}{}
		}
	}

	for i := 0; i < 9; i++ {
		for k := range flag {
			delete(flag, k)
		}
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			if board[j][i] < '1' || board[j][i] > '9' {
				return false
			}
			if _, ok := flag[board[j][i]]; ok {
				return false
			}
			flag[board[j][i]] = struct{}{}
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for k := range flag {
				delete(flag, k)
			}
			for k := i; k < i + 3; k++ {
				for l := j; l < j + 3; l++ {
					if board[k][l] == '.' {
						continue
					}
					if board[k][l] < '1' || board[k][l] > '9' {
						return false
					}
					if _, ok := flag[board[k][l]]; ok {

						return false
					}
					flag[board[k][l]] = struct{}{}
				}
			}
		}
	}
	return true
}

func main() {
	boardA := [][]byte{
		{'8','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	}
	fmt.Println(isValidSudoku(boardA))

	boardB := [][]byte{
		{'.','.','4','.','.','.','6','3','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'5','.','.','.','.','.','.','9','3'},
		{'.','.','.','5','6','.','.','.','.'},
		{'4','.','3','.','.','.','.','.','1'},
		{'.','.','.','7','.','.','.','.','.'},
		{'.','.','.','8','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
		{'.','.','.','.','.','.','.','.','.'},
	}
	fmt.Println(isValidSudoku(boardB))
}
