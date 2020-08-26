package main

import (
	"fmt"
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
*/

/*
错误想法，双指针可以解决，左右两个指针指向了数组的首尾，然后遍历计算，如果左指针为0，则交换左右指针值，然后移动指针。
第一次考虑错了，想的是互换，读题没读对，应该是将0移到最后
*/
func moveZeroes1(nums []int) {
	left, right := 0, len(nums)-1

	for left <= right {
		if nums[right] == 0 {
			right--
			continue
		}
		if nums[left] == 0 {
			nums = append(nums[:left], nums[left+1:]...)
			nums = append(nums, 0)
			continue
		}
		left++
	}
}

// 类似于创建新数组，只不过是在原数组操作。创建一个int类型存储新数组的长度，然后遍历原数组，如果不为零则进行赋值，如果为0则跳过赋值，newlen作为新的下标，记录的应该是有值的
/*
这个解法，通过新建一个下标用于重新赋值不为0的数
*/
func moveZeroes(nums []int) {
	newlen := 0

	for _, v := range nums {
		if v != 0 {
			nums[newlen] = v
			newlen++
		}
	}
	for ; newlen < len(nums); newlen++ {
		nums[newlen] = 0
	}
}

func Test_T283(t *testing.T) {
	nums := []int{0, 0, 1}
	moveZeroes(nums)
	fmt.Println(nums)
}
