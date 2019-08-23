package test

import (
	"fmt"
	"testing"
)

func Test_sliceRange(t *testing.T) {
	nums := []int{1, 2, 5, 4, 5}
	num := 0
	for k, v := range nums {
		if v == 5 {
			nums = append(nums[:k], nums[k+1:]...)
		}
		fmt.Println(num, v)
		fmt.Println(nums)
	}
}
