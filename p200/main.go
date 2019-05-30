package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != '0' {
				dfs(i, j, grid)
				count += 1
			}
		}
	}
	return count
}

func dfs(i, j int, grid[][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	// left
	dfs(i, j - 1, grid)
	// right
	dfs(i, j + 1, grid)
	// top
	dfs(i - 1, j, grid)
	// down
	dfs(i + 1, j, grid)
}

func main() {
	islands := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println(numIslands(islands))
}
