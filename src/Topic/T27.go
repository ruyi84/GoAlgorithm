package main

func removeElement(nums []int, val int) int {
	for k, v := range nums {
		if v == val {
			nums = append(nums[:k], nums[k+1:])
		}
	}
	return len(nums)
}
