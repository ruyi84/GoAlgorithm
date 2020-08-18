package main

import (
	"fmt"
	"testing"
)

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
初始想法，第一次删除一个数字时，删除的数字的下一个数字需要前移1，之后的数字也需要前移1，再次删除需要前移2
*/
func removeElement(nums []int, val int) int {
	newLen := 0
	for k := range nums {
		num := nums[k]
		if num != val {
			nums[newLen] = num
			newLen++
		}

	}
	return newLen
}

func removeElement2(nums []int, val int) int {
	i := 0
	n := len(nums)

	for i < n {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}

func Test_T027(t *testing.T) {
	nums := []int{3, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
	fmt.Println(nums)
}
