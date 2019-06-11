package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	result := make([]int, 0)
	if input == "" {
		return result
	}

	for i := 0; i < len(input); i++ {
		if '0' <= input[i] && '9' >= input[i] {
			continue
		}

		leftResult := diffWaysToCompute(input[:i])
		rightResult := diffWaysToCompute(input[i+1:])

		for _, left := range leftResult {
			for _, right := range rightResult {
				var rs int
				switch input[i] {
				case '-':
					rs = left - right
				case '*':
					rs = left * right
				default:
					rs = left + right
				}
				result = append(result, rs)
			}
		}
	}

	if len(result) == 0 {
		v, _ := strconv.Atoi(input)
		result = append(result, v)
	}
	return result
}


func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
