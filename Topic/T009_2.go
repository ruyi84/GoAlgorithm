package main

import "fmt"

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	y := x
	nums := 0
	for {
		nums = (x % 10) + nums*10
		x = x / 10
		if x == 0 {
			break
		}
	}
	if nums == y {
		return true
	}
	return false
}

func main() {
	i := 12321
	fmt.Println(isPalindrome2(i))
}
