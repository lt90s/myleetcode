package main

import "fmt"

func longestPalindrome(s string) string {
	longestI, longestJ := 0, 0
	for i := 0; i < len(s); i++ {
		if len(s) - i <= longestJ - longestI {
			break
		}
		for j := i + longestJ - longestI; j < len(s); j++ {
			if !isPalindrome(s, i, j) {
				continue
			}
			longestI = i
			longestJ = j + 1
		}
	}
	return s[longestI:longestJ]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome(""))
}