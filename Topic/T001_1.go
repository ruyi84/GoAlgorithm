package main

/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum2(nums []int, target int) []int {
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
