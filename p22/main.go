package main

import (
	"fmt"
	"strings"
)

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	result := []string{}
	part := make([]string, 2*n)
	doGenerate(n, 0, 0, part, &result)
	return result
}

func doGenerate(left, right, index int, part []string, result *[]string) {
	if left == 0 {
		for index < len(part) {
			part[index] = ")"
			index++
		}
		*result = append(*result, strings.Join(part, ""))
		return
	}

	part[index] = "("
	doGenerate(left-1, right+1, index+1, part, result)

	if right > 0 {
		part[index] = ")"
		doGenerate(left, right-1, index+1, part, result)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
