package main

import "fmt"

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
*/
func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			return i
		}
	}

	return i
}

func main() {

	arr := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(arr, 7))
}
