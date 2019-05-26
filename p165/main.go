package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")

	i := 0
	for ; i < len(parts1) && i < len(parts2); i++ {
		v1, _ := strconv.Atoi(parts1[i])
		v2, _ := strconv.Atoi(parts2[i])
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}

	if len(parts1) == len(parts2) {
		return 0
	}

	for i < len(parts1) {
		v, _ := strconv.Atoi(parts1[i])
		if v != 0 {
			return 1
		}
		i += 1
	}

	for i < len(parts2) {
		v, _ := strconv.Atoi(parts2[i])
		if v != 0 {
			return -1
		}
		i += 1
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1.0"))
	fmt.Println(compareVersion("1.1", "1.01"))
	fmt.Println(compareVersion("00001.1", "1.01"))
	fmt.Println(compareVersion("1.0", "1.0.0.0"))
}
