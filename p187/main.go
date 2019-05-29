package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	if length < 11 {
		return nil
	}
	countMap := make(map[string]int, length - 10)
	for i := 0; i <= length - 10; i++ {
		sub := s[i:i+10]
		if _, ok := countMap[sub]; !ok {
			countMap[sub] = 1
		} else {
			countMap[sub] += 1
		}
	}

	result := make([]string, 0, 16)
	for sub, count := range countMap {
		if count > 1 {
			result = append(result, sub)
		}
	}
	return result
}

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}
