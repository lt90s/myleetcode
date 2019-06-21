package main

import "fmt"

var letterMapping = map[uint8]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	result := []string{}
	recursive(digits, 0, "", &result)
	return result
}

func recursive(digits string, index int, s string, resultP *[]string) {
	if index >= len(digits) {
		*resultP = append(*resultP, s)
		return
	}

	for _, letter := range letterMapping[digits[index]] {
		recursive(digits, index+1, s+string(letter), resultP)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
