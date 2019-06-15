package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxArea(height []int) int {
	i, j := 0, len(height) - 1
	area := 0

	for i < j {
		tmp := minInt(height[i], height[j]) * (j - i)
		area = maxInt(area, tmp)

		if height[i] < height[j] {
			i += 1
		} else {
			j -= 1
		}
	}

	return area
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}