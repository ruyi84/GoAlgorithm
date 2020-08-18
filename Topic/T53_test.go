package main

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/
func maxSubArray(nums []int) int {
	maxNum := nums[0]
	for i := range nums {
		num := nums[i]
		if num > maxNum {
			maxNum = num
		}
		for j := i + 1; j < len(nums); j++ {
			num += nums[j]
			if num > maxNum {
				maxNum = num
			}
		}
	}
	return maxNum
}

func maxSubArray2(nums []int) int {
	l := len(nums)
	max := nums[0]

	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func Test_T53(t *testing.T) {
	fmt.Println(maxSubArray([]int{-1}))
}
