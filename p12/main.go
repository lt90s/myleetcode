package main

import "fmt"

var (
	symbols = [][]string{
		{"I", "V"},
		{"X", "L"},
		{"C", "D"},
		{"M", "-"},
	}
)

func intToRoman(num int) string {
	div, rs := 1000, ""
	for i := 3; i >= 0; i-- {
		tmp := num / div
		if tmp == 0 {
			div /= 10
			continue
		}
		rs += convert(tmp, i)
		num -= tmp * div
		div /= 10
	}
	return rs
}

func convert(num int, off int) string {
	var x string
	switch {
	case num == 4:
		x = symbols[off][0] + symbols[off][1]
	case num == 9:
		x = symbols[off][0] + symbols[off+1][0]
	case num < 5:
		for i := 0; i < num; i++ {
			x += symbols[off][0]
		}
	default:
		x = symbols[off][1]
		for i := 5; i < num; i++ {
			x += symbols[off][0]
		}
	}
	return x
}

func main() {
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(900))
}
