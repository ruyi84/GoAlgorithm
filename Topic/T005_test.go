package main

func longestPalindrome_1(s string) string {
	len := len(s)

	if len < 2 {
		return s
	}

	max := 1
	begin := 0

	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if (j-i+1) > max && validPalindromic(s, i, j) {
				max = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+max]
}

func validPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
