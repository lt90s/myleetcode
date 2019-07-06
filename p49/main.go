package main

import (
	"fmt"
	"sort"
)

type byteSlice []byte

func (b byteSlice) Len() int {
	return len(b)
}

func (b byteSlice) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b byteSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	for _, str := range strs {
		bs := byteSlice(str)
		sort.Sort(bs)
		s := string(bs)
		if _, ok := group[s]; !ok {
			group[s] = make([]string, 0, 1)
		}
		group[s] = append(group[s], str)
	}

	rs := make([][]string, 0, len(group))
	for _, g := range group {
		rs = append(rs, g)
	}
	return rs
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
