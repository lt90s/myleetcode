package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	bNum1 := []byte(num1)
	bNum2 := []byte(num2)
	maxSize := len(bNum1) * len(bNum2) + 1
	res := make([]byte, maxSize)
	tmp := make([]byte, len(bNum1) + 1)

	step := 0

	for i := len(bNum2) - 1; i >= 0; i-- {
		m := bNum2[i] - '0'
		index := 0
		carry := 0
		for j := len(bNum1) - 1; j >= 0; j-- {
			n := bNum1[j] - '0'
			r := int(m) * int(n) + carry
			tmp[index] = byte(r % 10)
			carry = r / 10
			index++
		}
		tmp[index] = byte(carry)
		k, carry := 0, 0
		for ; k < len(tmp); k++ {
			r := int(res[step+k] + tmp[k]) + carry
			res[step+k] = byte(r % 10)
			carry = r / 10
		}
		step++
	}
	i, j := 0, len(res) - 1
	for res[j] == 0 {
		j--
	}
	res = res[:j+1]

	for i <= j {
		res[i], res[j] = res[j] + '0', res[i] + '0'
		i++
		j--
	}
	return string(res)
}

func main() {
	fmt.Println(multiply("9", "99"))
}
