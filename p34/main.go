package main

import (
	"fmt"
	"sort"
)


func searchRange(nums []int, target int) []int {
	lower := sort.SearchInts(nums, target)
	if lower >= len(nums) || nums[lower] != target {
		return []int{-1, -1}
	}

	upper := sort.SearchInts(nums, target + 1)
	if upper - 1 < 0 {
		return []int{-1, -1}
	}

	return []int{lower, upper - 1}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 5))
}
