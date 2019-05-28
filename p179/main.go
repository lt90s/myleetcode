package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type seqNums []string

func (s seqNums) Len() int {
	return len(s)
}

func (s seqNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s seqNums) Less(i, j int) bool {
	return s[i] + s[j] < s[j] + s[i]
}

func largestNumber(nums []int) string {
	sum := 0
	s := make(seqNums, 0, len(nums))
	for _, num := range nums {
		if sum == 0 {
			sum += num
		}
		s = append(s, strconv.Itoa(num))
	}
	if sum == 0 {
		return "0"
	}
	sort.Sort(sort.Reverse(s))
	fmt.Println(s)
	return strings.Join(s, "")
}

func main() {
	nums := []int{824,938,1399,5607,6973,5703,9609,4398,8247}
	fmt.Println(largestNumber(nums))
}
