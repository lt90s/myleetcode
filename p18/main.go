package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var a, b, c *int
	result := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if a != nil && *a == nums[i] {
			continue
		}
		a = &nums[i]
		b, c = nil, nil
		for j := i + 1; j < len(nums)-2; j++ {
			if b != nil && *b == nums[j] {
				continue
			}
			b = &nums[j]
			c = nil
			for k := j + 1; k < len(nums)-1; k++ {
				if c != nil && *c == nums[k] {
					continue
				}
				c = &nums[k]
				d := target - *a - *b - *c
				index := sort.SearchInts(nums[k+1:], d)
				if index >= len(nums[k+1:]) {
					continue
				}

				if nums[k+1+index] == d {
					result = append(result, []int{*a, *b, *c, d})
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
}
