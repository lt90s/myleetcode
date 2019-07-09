package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}

	for j := 0; j < n; j++ {
		dp[m-1][j] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}

	return dp[0][0]
}

func main() {
	fmt.Println(uniquePaths(23, 12))
	fmt.Println(uniquePaths(7, 3))
}
