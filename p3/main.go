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

func lengthOfLongestSubstring_On(s string) int {
	indexes := [256]int{}
	rs, cursor := 0, 0
	for i := 0; i < len(s); i++ {
		// repeat
		if indexes[s[i]] != 0 && indexes[s[i]] >= cursor {
			cursor = indexes[s[i]]
		}
		indexes[s[i]] = i + 1
		if i - cursor + 1 > rs {
			rs = i - cursor + 1
		}
	}
	return rs
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))

	fmt.Println(lengthOfLongestSubstring_On("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring_On("bbbbb"))
	fmt.Println(lengthOfLongestSubstring_On("pwwkew"))
	fmt.Println(lengthOfLongestSubstring_On(" "))
}
