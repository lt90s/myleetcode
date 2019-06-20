package main

import (
	"fmt"
	"math"
	"sort"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var minDiff int64 = math.MaxInt64
	var result int

	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return sum
			}

			diff := absInt(sum - target)
			if int64(diff) < minDiff {
				minDiff = int64(diff)
				result = sum
			}

			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
