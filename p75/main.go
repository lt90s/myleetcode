package main

import "fmt"

func sortColors(nums []int) {
	i, j, k := 0, 0, 0
	for _, num := range nums {
		nums[i] = 2
		i++
		if num < 2 {
			nums[j] = 1
			j++
		}
		if num < 1 {
			nums[k] = 0
			k++
		}

	}
}

func main() {
	nums := []int{0, 0, 2}
	sortColors(nums)
	fmt.Println(nums)
}
