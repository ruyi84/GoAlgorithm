package main

func twoSum(nums []int, target int) []int {
	i := make([]int, 2, 2)
	for k1, v1 := range nums {
		for k2, v2 := range nums {
			if target == (v1+v2) && k1 != k2 {
				i[0] = k1
				i[1] = k2
				return i
			}
		}
	}
	return nil
}
