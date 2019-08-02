package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	i, j, left := 0, 0, 0
	for i < len(gas) {
		left = 0
		for j = i; j < len(gas); j++ {
			left += gas[j] - cost[j]
			if left < 0 {
				break
			}
		}
		if j >= len(gas) {
			break
		}
		i = j + 1
	}

	for j := 0; j < i; j++ {
		left += gas[j] - cost[j]
	}

	if left < 0 {
		return -1
	} else {
		return i
	}
}

func main() {
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
