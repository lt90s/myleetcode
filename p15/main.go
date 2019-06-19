package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}

	var lastA, lastB *int
	for i := 0; i < len(nums); i++ {
		if lastA != nil && nums[i] == *lastA {
			continue
		}
		lastB = nil
		for j := i + 1; j < len(nums); j++ {
			if lastB != nil && nums[j] == *lastB {
				continue
			}
			sum := nums[i] + nums[j]
			index := sort.SearchInts(nums[j+1:], -sum)
			if index >= len(nums[j+1:]) || nums[j+1+index] != -sum {
				continue
			}

			lastA = &nums[i]
			lastB = &nums[j]
			result = append(result, []int{nums[i], nums[j], -sum})
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}