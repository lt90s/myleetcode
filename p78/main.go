package main

import "fmt"

func subsets(nums []int) [][]int {
	n := len(nums)
	rv := make([][]int, 0)
	tmp := make([]int, 0, n)
	for i := 0; i < (1 << uint(n)); i++ {
		tmp = tmp[:0]
		k := 0
		for j := i; j != 0; j = j >> 1 {
			if j & 1 == 1 {
				tmp = append(tmp, nums[k])
			}
			k++
		}
		ttmp := make([]int, len(tmp))
		copy(ttmp, tmp)
		rv = append(rv, ttmp)
	}
	return rv
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}