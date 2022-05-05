package main

import "testing"

func Test704(t *testing.T) {
	t.Log(search2([]int{2, 5}, 5))
}

// 直接遍历法
func search1(nums []int, target int) int {
	for index, num := range nums {
		if num == target {
			return index
		}
	}

	return -1
}

// 尝试二分查找，因为nums是有序的数组
func search2(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func search3(nums []int, target int) int {
	mid, left, right := len(nums)/2, 0, len(nums)-1

	for left <= right {
		if nums[mid] == target{
			return mid
		}

		if nums[mid] > target{
			right = mid-1

		}else {
			left = mid+1
		}

		mid = (left+right) /2

	}

	return -1
}
