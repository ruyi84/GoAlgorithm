package main

import "fmt"

func longestPalindrome(s string) string {

	for k, v := range s {
		for i := 0; i < len(s); i++ {
			if v == s[len(s)-i] {

			}
		}
	}
	return s
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}

/*
1、从字符串两侧像中取，如果出现不同则左侧下表右移
2、
*/
