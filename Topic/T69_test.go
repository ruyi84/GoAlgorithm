package main

import (
	"testing"
)

func mySqrt(x int) int {
	left, right := 0, x

	result := 0
	for left <= right {
		mid := left + (right-left)/2

		num := mid * mid
		//if num == x {
		//	return mid
		//}

		if num > x {
			right = mid - 1
		} else {
			result = mid
			left = mid + 1
		}
	}

	return result
}

func TestT69(t *testing.T) {
	t.Log(mySqrt(2))
	t.Log(mySqrt(4))
	t.Log(mySqrt(8))
}
