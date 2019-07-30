package main

import "fmt"

func partition(s string) [][]string {
	current := []string{}
	rv := [][]string{}
	if len(s) == 0 {
		return nil
	}
	dfs(0, s, current, &rv)
	return rv
}

func dfs(idx int, s string, current []string, rv *[][]string) {
	if idx == len(s) {
		tmp := make([]string, len(current))
		copy(tmp, current)
		*rv = append(*rv, tmp)
		return
	}

	if idx > len(s) {
		return
	}

	for i := idx; i < len(s); i++ {
		if !valid(s[idx : i+1]) {
			continue
		}
		dfs(i+1, s, append(current, s[idx:i+1]), rv)
	}
}

func valid(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(partition("aabba"))
}
