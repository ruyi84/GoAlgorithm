package main

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。
*/

/*
这个题目是要查看一个数组中是否存在两个相等数值的下表绝对值小于k
	方法1，双指针，前一个指针确定后，遍历该指针后的k个值是否存在相同值
	方法2，创建一个map，key为数组中的值，value为对应的下标，遍历数组时，如果该值存在，则计算绝对值差，如果不存在则存入map
*/

//func1 双指针，前一个指针确定后，遍历该指针后的k个值是否存在相同值,976ms,5.1MB
func containsNearbyDuplicate1(nums []int, k int) bool {
	for key := range nums {
		num1 := nums[key]
		for i := key + 1; i < len(nums) && i <= key+k; i++ {
			if num1 == nums[i] {
				return true
			}
		}

	}
	return false
}

// func2 创建一个map，key为数组中的值，value为对应的下标，遍历数组时，如果该值存在，则计算绝对值差，如果不存在则存入map 20ms,7.8MB
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 1 {
		return false
	}
	m := make(map[int]int)
	m[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num == nums[0] && i <= k {
			return true
		}
		if m[num] != 0 && i-m[num] <= k {
			return true
		}
		m[num] = i
	}

	return false
}

// func2改进,原本担心的是出现下标为0相同时会出现误解，但是可以通过j, ok := m[num]的方式进行取值，可以判断是否存在值
func containsNearbyDuplicate2(nums []int, k int) bool {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[num]; ok && i-j <= k {
			return true
		}
		m[num] = i
	}
	return false
}

func Test_T219(t *testing.T) {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}
