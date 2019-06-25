package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	num := 999
	for i := 0; i < len(nums); i++ {
		if i > 0 && num == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		num = nums[i]
	}
	return len(nums)
}

func main() {
	a := []int{1, 1, 2, 3}
	fmt.Println(removeDuplicates(a))
}
