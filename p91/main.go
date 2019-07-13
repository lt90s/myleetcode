package main

import "fmt"

//dp[i] = dp[i+1] + dp[i+2]?
func numDecodings(s string) int {
	length := len(s)
	dp := make([]int, length+1)
	dp[length] = 1

	if s[length - 1] != '0' {
		dp[length-1] = 1
	}

	for i := length - 2; i >= 0; i-- {
		if s[i] =='0' {
			dp[i] = 0
		} else if s[i] == '1' {
			dp[i] = dp[i+1] + dp[i+2]
		} else if s[i] == '2' {
			if s[i+1] >= '0' && s[i+1] <= '6' {
				dp[i] = dp[i+1] + dp[i+2]
			} else {
				dp[i] = dp[i+1]
			}
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}

func main() {
	fmt.Println(numDecodings("117262"))
}
