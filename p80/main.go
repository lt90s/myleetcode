package main

import "fmt"

func removeDuplicates(nums []int) int {
	p1, p2 := 0, 1
	if len(nums) <= 2 {
		return len(nums)
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[p1] {
			if p2-p1 >= 2 {
				continue
			}
			nums[p2] = nums[i]
			p2++
		} else {
			nums[p2] = nums[i]
			p1 = p2
			p2++
		}
	}
	return p2
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	n := removeDuplicates(nums)
	fmt.Println(nums[:n])
}
