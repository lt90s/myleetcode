package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	i, j := 0, 0
	tmp := 0
	minSize := len(nums) + 1
	for ; j < len(nums); j++ {
		tmp += nums[j]
		if tmp < s {
			continue
		}

		for tmp >= s {
			if j - i + 1 < minSize {
				minSize = j - i + 1
			}
			tmp -= nums[i]
			i += 1
		}
	}

	if minSize == len(nums) + 1 {
		return 0
	}
	return minSize
}


func main() {
	fmt.Println(minSubArrayLen(7, []int{2,1,1,2,1,0}))
}