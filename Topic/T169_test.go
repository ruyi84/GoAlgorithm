package main

import (
	"fmt"
	"sort"
	"testing"
)

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 哈希表计数
func majorityElement1(nums []int) int {
	length := len(nums)
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
		if m[v] > length/2 {
			return v
		}
	}
	return 0
}

// 排序，已知必然存在众数，则排序后的中间值必定为众数
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*
摩尔投票分为两个阶段
	1.对抗阶段：分属两个候选人的票数进行两两对抗抵消
	2.计数阶段：计算对抗结果中最后留下的候选人票数是否有效

因题目中已知给定的数组总存在多数元素，且多数元素占总数的1/2，故多数元素的数量可以抵消所有的数量
*/
func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			major = v
			count++
			continue
		}
		if v != major {
			count--
			continue
		}
		count++
	}
	return major
}

func Test_T169(t *testing.T) {
	fmt.Println(majorityElement([]int{6, 5, 5}))
}
