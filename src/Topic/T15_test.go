package main

import (
	"fmt"
	"testing"
)

func threeSum(nums []int) [][]int {
	m1 := make(map[int]int)

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := i + 2; k < len(nums)-1; k++ {
				if (nums[i] + nums[j] + nums[k]) == 0 {

				}

			}
		}
	}
	return m1
}

func Test15(t *testing.T) {
	i := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(i))
}
