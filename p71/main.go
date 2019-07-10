package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) == 0 || parts[0] != "" {
		return ""
	}

	paths := []string{}

	for i := 1; i < len(parts); i++ {
		switch parts[i] {
		case ".", "":
			continue
		case "..":
			if len(paths) == 0 {
				continue
			}
			paths = paths[:len(paths)-1]
		default:
			paths = append(paths, parts[i])
		}
	}
	return "/" + strings.Join(paths, "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
