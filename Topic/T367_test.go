package main

import (
	"fmt"
	"testing"
)

// 判断num是否是完全平方数
// 这是二分查找的题目，所以直接使用二分法
func isPerfectSquare(num int) bool {
	left, right := 0, num

	for left <= right {
		mid := (left + right + 1) / 2
		fmt.Println(mid, left, right)
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

func TestT367(t *testing.T) {
	t.Log(isPerfectSquare(14))
}
