package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	current := []int{}
	res := [][]int{}

	if target == 0{
		return nil
	}
	sort.Ints(candidates)
	do(candidates, 0, target, current, &res)
	return res
}

func do(candidates []int, sum, target int, current []int, res *[][]int) {
	if sum == target {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*res = append(*res, tmp)
		return
	}
	if len(candidates) == 0  || sum > target{
		return
	}
	i := 1
	for ; i < len(candidates); i++ {
		if candidates[i] != candidates[i-1] {
			break
		}
	}

	for j := i; j >= 0; j-- {
		do(candidates[i:], sum, target, current, res)
		current = append(current, candidates[0])
		sum += candidates[0]
	}
}


func main() {
	nums := []int{10,1,2,7,6,1,5}
	fmt.Println(combinationSum2(nums, 8))
}