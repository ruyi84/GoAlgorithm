package DynamicProgramming

import (
	"fmt"
	"testing"
)

func Test_CutWood(t *testing.T) {
	ints := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	fmt.Println(cut(ints, 2))
	fmt.Println(cutMemo(ints))
}

func cut(p []int, n int) int {
	if n == 0 {
		return n
	}

	q := 0
	for i := 1; i <= n; i++ {
		q = max(q, p[i-1]+cut(p, n-i))
	}
	return q

}

func cutMemo(p []int) int {
	ints := make([]int, len(p)+1)
	for i := 0; i < len(p); i++ {
		ints[i] = -1
	}
	return cut1(p, len(p), ints)
}

func cut1(p []int, n int, r []int) int {
	q := -1
	if r[n] >= 0 {
		return r[n]
	}

	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i++ {
			q = max(q, cut1(p, n-i, r)+p[i-1])
		}
		r[n] = q
	}
	return q
}

func buttom_up_cut(p []int) int {
	ints := make([]int, len(p)+1)

	for i := 1; i <= len(p); i++ {
		q := -1
		for j := 1; j <= i; j++ {
			q = max(q, p[j-1]+ints[i-j])
		}
		ints[i] = q
	}
	return ints[len(p)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
