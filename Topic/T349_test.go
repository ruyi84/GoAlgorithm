package main

import (
	"sort"
	"testing"
)

//给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

// 双重for循环
// 因为内部可能存在相同的值，借助map进行去重
func intersection1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	resultMap := make(map[int]struct{})
	for _, v := range nums1 {
		for _, v2 := range nums2 {
			if v == v2 {
				resultMap[v] = struct{}{}
			}
		}
	}

	var result []int
	for k, _ := range resultMap {
		result = append(result, k)
	}

	return result
}

// 因为训练目标是二分法，所以使用二分法
// 因为改题目中入参为无序切片，所以需要先进行排序
// 开始遇到一个困难，就是会存在（7+8）/ 2 = 7的死循环问题，后来操作时为right和left分别+-1，解决奇数问题
func intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	resultMap := make(map[int]struct{})

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	for _, v2 := range nums1 {
		nums2Mid := len(nums2) / 2
		left, right := 0, len(nums2)-1
		for left <= right {
			if v2 == nums2[nums2Mid] {
				resultMap[v2] = struct{}{}
				break
			}

			if v2 > nums2[nums2Mid] {
				left = nums2Mid + 1
				nums2Mid = (nums2Mid + right + 1) / 2
				continue
			} else {
				right = nums2Mid - 1
				nums2Mid = (nums2Mid + left - 1) / 2
				continue
			}
		}
	}

	var result []int
	for k := range resultMap {
		result = append(result, k)
	}

	return result
}

func TestT349(t *testing.T) {
	t.Log(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	t.Log(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	t.Log(intersection([]int{1}, []int{1, 2}))
	t.Log(intersection([]int{1}, []int{}))
	t.Log(intersection([]int{4, 7, 9, 7, 6, 7}, []int{5, 0, 0, 6, 1, 6, 2, 2, 4}))
	t.Log(intersection([]int{0, 5, 8, 7, 2, 9, 7, 5}, []int{1, 4, 8, 9}))
}
