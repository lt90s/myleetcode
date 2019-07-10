package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}

	j := n - 1
	for ; j >= 0; j-- {
		if obstacleGrid[m-1][j] == 0 {
			dp[m-1][j] = 1
		} else {
			break
		}
	}
	for ; j >= 0; j-- {
		dp[m-1][j] = 0
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}

func main() {
	grid := [][]int{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(grid))
}
