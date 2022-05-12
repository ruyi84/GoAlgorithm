package main

import "testing"

func searchRange(nums []int, target int) []int {
	return []int{left(nums,target),right(nums,target)}
}

func left(nums []int, target int) int {
	left,right := 0,len(nums)-1

	result := 0

	for left < right{
		mid := left +(right -left ) /2

		result = mid

		if nums[mid] >= target{
			right = mid
		}else{
			left = mid +1
		}
	}

	if nums[result] == target{
		return result
	}

	return -1
}


func right(nums []int, target int) int {
	left,right := 0,len(nums)-1

	result := 0

	for left < right{
		mid := left +(right -left ) /2

		result = mid

		if nums[mid] <= target{
			left = mid +1
		}else{
			right = mid
		}
	}

	if nums[result] == target{
		return result
	}

	return -1
}

func TestT34(t *testing.T) {
	t.Log(searchRange([]int{5,7,7,8,8,10},8))
}
