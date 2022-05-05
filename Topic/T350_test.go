package main

import (
	"fmt"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	resultMap := make(map[int]int)
	for _, v := range nums1 {
		resultMap[v]++
	}

	fmt.Println(resultMap)
	var result []int
	for _, v := range nums2 {
		if nums, ok := resultMap[v]; ok && nums > 0 {
			resultMap[v]--
			result = append(result, v)
		}
	}

	return result
}

func TestT350(t *testing.T) {
	//t.Log(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	//t.Log(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
	//t.Log(intersect([]int{1}, []int{1}))
	t.Log(intersect([]int{4,9,5}, []int{9,4,9,8,4}))
}
