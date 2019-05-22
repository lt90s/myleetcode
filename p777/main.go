package main

import "fmt"

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	i, j, length := 0, 0, len(start)
	for i < length {
		for i < length && start[i] == 'X' {
			i++
		}

		for j < length && end[j] == 'X' {
			j++
		}

		if i >= length && j >= length {
			return true
		}

		if start[i] != end[j] {
			return false
		}

		if start[i] == 'L'{
			if i < j {
				return false
			}
		} else if i > j {
			return false
		}
		i += 1
		j += 1
	}
	return true
}

func main() {
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))
}
