package main

import "fmt"

func singleNumber(nums []int) int {
	count := make([]int, 64)
	for _, num := range nums {
		for i := 0; i < 64; i++ {
			if num&(1<<uint(i)) != 0 {
				count[i]++
			}
		}
	}
	rv := 0
	for i, c := range count {
		rv |= (c % 3) << uint(i)
	}
	return rv
}

func main() {
	nums := []int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}
	fmt.Println(singleNumber(nums))
}
