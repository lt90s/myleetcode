package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	flags := make(map[byte]struct{})
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for key := range flags {
			delete(flags, key)
		}
		for j := i; j < len(s); j++ {
			if _, ok := flags[s[j]]; ok {
				break
			}
			flags[s[j]] = struct{}{}
		}
		if maxLen < len(flags) {
			maxLen = len(flags)
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
