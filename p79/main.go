package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	found := false
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			search(board, word, i, j, 0, visited, &found)
			if found {
				return true
			}
		}
	}
	return false
}

func search(board[][]byte, word string, i, j, k int, visited [][]bool, found *bool) {
	if *found {
		return
	}

	if k == len(word) {
		*found = true
		return
	}

	if i >= len(board) || i < 0 || j < 0 || j >= len(board[0]) {
		return
	}

	if visited[i][j] {
		return
	}

	if board[i][j] != word[k] {
		return
	}

	visited[i][j] = true
	search(board, word, i+1, j, k+1, visited, found)
	search(board, word, i-1, j, k+1, visited, found)
	search(board, word, i, j+1, k+1, visited, found)
	search(board, word, i, j-1, k+1, visited, found)
	visited[i][j] = false
}

func main() {
	fmt.Println(exist([][]byte{{'a',}, {'b',}}, "ba"))
}