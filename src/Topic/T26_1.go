package main

import "fmt"

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) <= 1 {
		return len(nums)
	}
	n := len(nums)
	slow, fast := 0, 0
	for fast < n {
		for fast < n-1 && nums[fast] == nums[fast+1] {
			fast++
		}
		nums[slow] = nums[fast]
		slow += 1
		fast += 1

	}
	return slow
}

func main() {
	a := []int{1, 1, 2}
	fmt.Println(removeDuplicates2(a))
}
