package main

import (
	"fmt"
	"strconv"
)

func rangeBitwiseAnd(m int, n int) int {
	ms := fmt.Sprintf("%b", m)
	ns := fmt.Sprintf("%b", n)

	if len(ms) != len(ns) {
		return 0
	}

	tmp := make([]byte, len(ms))
	i := 0
	for i < len(ms) {
		if ms[i] != ns[i] {
			break
		}
		tmp[i] = ms[i]
		i += 1
	}
	for i < len(ms) {
		tmp[i] = '0'
		i += 1
	}
	rs, _ := strconv.ParseInt(string(tmp), 2, 32)
	return int(rs)
}

func main() {
	fmt.Println(rangeBitwiseAnd(4, 7))
}