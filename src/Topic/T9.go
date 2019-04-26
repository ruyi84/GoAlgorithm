package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := make([]int, 10)
	i := 0
	b := true
	for {
		if x == 0 {
			break
		}
		s[i] = x % 10
		x /= 10
		i += 1
	}
	for n := 0; n < i/2; n++ {
		if s[n] != s[i-1-n] {
			b = false
		}
	}

	return b
}

func main() {
	i := 12321
	fmt.Println(isPalindrome(i))
}
