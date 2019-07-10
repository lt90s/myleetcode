package main

import "fmt"

func combine(n int, k int) [][]int {
	current := []int{}
	rv := [][]int{}
	if n <= 0 || k <= 0 {
		return nil
	}
	doCombine(1, n, k, current, &rv)
	return rv
}

func doCombine(i, n, k int, current []int, rv *[][]int) {
	if k == 0 {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*rv = append(*rv, tmp)
		return
	}

	if i > n || n - i + 1 < k {
		return
	}

	doCombine(i+1, n, k - 1, append(current, i), rv)
	doCombine(i+1, n, k, current, rv)
}

func main() {
	fmt.Println(combine(4, 2))
}