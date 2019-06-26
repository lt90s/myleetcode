package main

import (
	"fmt"
	"sort"
)

func reverse(nums []int) {
	i, j := 0, len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func nextPermutation(nums []int)  {
	i := len(nums) -1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i <= 0 {
		reverse(nums)
		return
	}

	decrease := nums[i:]
	index := sort.Search(len(decrease), func(j int) bool { return decrease[j] <= nums[i-1]})
	index = index - 1

	nums[i-1], decrease[index] = decrease[index], nums[i-1]
	reverse(decrease)
}

func display(nums []int) {
	if len(nums) == 0 {
		fmt.Println("[]")
	}
	fmt.Printf("%d", nums[0])
	for i := 1; i < len(nums); i++ {
		fmt.Printf(", %d", nums[i])
	}
	fmt.Printf("\n")
}

func permutationCount(n int) int {
	if n < 0 {
		return 0
	}
	rv := 1
	for n > 0 {
		rv *= n
		n--
	}
	return rv
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	display(nums)
	for i := 0; i < permutationCount(len(nums)); i++ {
		nextPermutation(nums)
		display(nums)
	}
}