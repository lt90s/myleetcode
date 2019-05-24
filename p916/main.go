package main

import "fmt"

func countWord(word string) []int {
	count := make([]int, 26)
	for _, w := range word {
		index := int(w - 'a')
		count[index] += 1
	}
	return count
}

func maxCountWords(words []string) []int {
	maxCount := make([]int, 26)
	for _, word := range words {
		count := countWord(word)
		for i := 0; i < 26; i++ {
			if count[i] > maxCount[i] {
				maxCount[i] = count[i]
			}
		}
	}
	return maxCount
}

func wordSubsets(A []string, B []string) []string {
	result := make([]string, 0)
	maxCount := maxCountWords(B)
	for _, word := range A {
		count := countWord(word)
		match := true
		for i := 0; i < 26; i++ {
			if maxCount[i] > count[i] {
				match = false
				break
			}
		}
		if match {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	A := []string{"amazon","apple","facebook","google","leetcode"}
	B := []string{"e","o"}
	fmt.Println(wordSubsets(A, B))
}
