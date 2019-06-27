package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] == target {
			return m
		} else if nums[m] > nums[l] {
			if nums[l] <= target && target <= nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else if nums[m] < nums[r] {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			l++
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1}
	fmt.Println(search(nums, 1))
	fmt.Println(search(nums, 3))
}
