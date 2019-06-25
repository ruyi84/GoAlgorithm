package main

import "fmt"

func romanToInt(s string) int {
	var nums = make(map[rune]int)
	nums['I'] = 1
	nums['V'] = 5
	nums['X'] = 10
	nums['L'] = 50
	nums['C'] = 100
	nums['D'] = 500
	nums['M'] = 1000

	num := 0
	n := 'I'
	for _, v := range s {
		if (nums[n] < nums[v]) && num != 0 {
			num = num - (nums[n] * 2)
		}
		num += nums[v]
		n = v
	}
	return num
}

func main() {
	s := "MCMXCIV"

	fmt.Println(romanToInt(s))
}
