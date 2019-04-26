package main

import "fmt"

func intToRoman(num int) string {

	var nums = make(map[int]rune)
	nums[1] = 'I'
	nums[5] = 'V'
	nums[10] = 'X'
	nums[50] = 'L'
	nums[100] = 'C'
	nums[500] = 'D'
	nums[1000] = 'M'

	var s string
	if num == 0 {
		return s
	}
	for {
		if num/1000 != 0 {
			s = s + nums[1000]
		}
	}
	return s
}

func main() {
	num := 3

	fmt.Println(romanToInt(num))
}
