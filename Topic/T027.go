package main

import "fmt"

func removeElement(nums []int, val int) int {
	for k, v := range nums {
		if v == val {
			if k+1 == len(nums) {
				nums = append(nums[:k])
				return len(nums)
			}
			nums = append(nums[:k], nums[k+1:]...)
		}
	}
	fmt.Println(nums)
	return len(nums)
}

func main() {
	nums := []int{3, 3}
	val := 3
	fmt.Println(removeElement(nums, val))

}
