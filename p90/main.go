package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	rv := [][]int{{}}
	if len(nums) == 0 {
		return rv
	}
	current := make([]int, len(nums))
	for i := 1; i <= len(nums); i++ {
		subsets(nums, 0, current, 0, i, &rv)
	}
	return rv
}

func subsets(nums []int, idx int, current []int, cIdx int, count int, rv *[][]int) {
	if cIdx >= count {
		tmp := make([]int, count)
		copy(tmp, current)
		*rv = append(*rv, tmp)
		return
	}

	if idx >= len(nums) {
		return
	}

	nextIdx := idx + 1
	for nextIdx < len(nums) {
		if nums[nextIdx] != nums[idx] {
			break
		}
		nextIdx++
	}

	subsets(nums, nextIdx, current, cIdx, count, rv)
	for idx < nextIdx && cIdx < count {
		current[cIdx] = nums[idx]
		idx++
		cIdx++
		subsets(nums, nextIdx, current, cIdx, count, rv)
	}
}

func main() {
	fmt.Println(subsetsWithDup([]int{4, 4, 1, 4}))
}
