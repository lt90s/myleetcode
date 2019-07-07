package main

import (
	"fmt"
	"sort"
)

type Intervals [][]int
func (itv Intervals) Len() int {
	return len(itv)
}

func (itv Intervals) Less(i, j int) bool {
	if itv[i][0] < itv[j][0] {
		return true
	}

	if itv[i][0] == itv[j][0] && itv[i][1] <= itv[j][1] {
		return true
	}

	return false
}

func (itv Intervals) Swap(i, j int) {
	itv[i], itv[j] = itv[j], itv[i]
}

func merge(intervals [][]int) [][]int {
	rv := [][]int{}

	if len(intervals) == 0 {
		return rv
	}

	itvs := Intervals(intervals)
	sort.Sort(itvs)

	itv := itvs[0]
	for i := 1; i < len(itvs); i++ {
		if itvs[i][0] > itv[1] {
			rv = append(rv, itv)
			itv = itvs[i]
		}
		if itv[1] < itvs[i][1] {
			itv[1] = itvs[i][1]
		}
	}
	rv = append(rv, itv)
	return rv
}

func main() {
	itvs := [][]int {
		{1,3},{2,6},{8,10},{15,18},
	}
	fmt.Println(merge(itvs))
}