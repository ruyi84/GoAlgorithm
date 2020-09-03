package main

/*
给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
*/

func thirdMax(nums []int) int {
	var MIN = -1 << 63
	a, b, c := MIN, MIN, MIN

	for _, n := range nums {
		if n == a || n == b || n == c {
			continue
		}

		if n > a {
			a, b, c = n, a, b
		} else if n > b {
			b, c = n, b
		} else if n > c {
			c = n
		}
	}

	if c == MIN {
		return a
	}
	return c
}
