package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	length := len(graph)
	visited := make([]bool, length)

	visited[0] = true
	result := make([][]int, 0)
	path := make([]int, length)
	path[0] = 0

	doDfs(graph, 0, 1, length - 1, path, &result, visited)
	return result
}

func doDfs(graph [][]int, cPoint, cLen, target int, path []int, result* [][]int, visited []bool) {
	if cPoint == target {
		tmp := make([]int, cLen)
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	for _, next := range graph[cPoint] {
		if visited[next] {
			continue
		}
		visited[next] = true
		path[cLen] = next
		doDfs(graph, next, cLen+1, target, path, result, visited)
		visited[next] = false
	}
}

func main() {
	graph := [][]int {
		{1, 2},
		{3},
		{3},
		{},
	}

	fmt.Println(allPathsSourceTarget(graph))
}
