package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	current := []int{}
	rs := [][]int{}
	do(candidates, target, current, &rs)
	return rs
}

func do(candidates []int, target int, current []int, rs *[][]int) {
	if target == 0 {
		if len(current) > 0 {
			tmp := make([]int, len(current))
			copy(tmp, current)
			*rs = append(*rs, tmp)
		}
		return
	}

	if len(candidates) == 0 {
		return
	}

	if candidates[0] > target {
		return
	}

	do(candidates, target-candidates[0], append(current, candidates[0]), rs)
	do(candidates[1:], target, current, rs)
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
