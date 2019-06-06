package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	leftIndex := -100
	result := make([]string, 0, 64)
	for i := 0; i < len(nums); i++ {
		if leftIndex == -100 {
			leftIndex = i
		} else {
			if nums[i - 1] + 1 != nums[i] { // not consecutive
				if leftIndex + 1 == i {
					result = append(result, strconv.Itoa(nums[leftIndex]))
				} else {
					result = append(result, fmt.Sprintf("%d->%d", nums[leftIndex], nums[i - 1]))
				}
				leftIndex = i
			} else if i == len(nums) - 1 {
				result = append(result, fmt.Sprintf("%d->%d", nums[leftIndex], nums[i]))
			}
		}
	}
	if leftIndex == len(nums) - 1 {
		result = append(result, strconv.Itoa(nums[leftIndex]))
	}
	return result
}

func main() {
	fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
	fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}
