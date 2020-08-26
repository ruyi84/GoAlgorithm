package main

import (
	"fmt"
	"testing"
)

/*
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
*/

/*
	方法1，排序后遍历，发现断层
*/

func missingNumber1(nums []int) int {
	sum := SumBylen(len(nums))

	for k := range nums {
		sum -= nums[k]
	}
	return sum
}

func SumBylen(len int) int {
	sum := 0
	for i := 0; i <= len; i++ {
		sum += i
	}
	return sum
}

// 已知，给定数组包含了一个序列，如果不缺少的话值应该是1到len的集，而实际缺少一个数字，且数组从0开始，则可以直接计算下标值+1累加为目标的和，然后和value进行差值计算，最后的结果就是缺少的值
func missingNumber2(nums []int) int {
	sum := 0

	for k := range nums {
		sum += k + 1
		sum -= nums[k]
	}
	return sum
}

// 利用了高斯算法，一个有序数集的和为头+尾的和*长度/2
func missingNumber(nums []int) int {
	len := len(nums)

	sum := (len + 1) * len / 2
	for k := range nums {
		sum -= nums[k]
	}
	return sum

}

func Test_T268(t *testing.T) {
	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums))
}
