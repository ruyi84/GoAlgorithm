package main

import (
	"fmt"
	"testing"
)

/*
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。



说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
如果nums1为空，则说明结果为nums2
双指针，倒叙遍历两个数组，大的放在nums1结束位置，当一个为空后遍历另一个结束就可以了
*/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i := len(nums1)
	for n != 0 && m != 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}
		i--
	}
	for !(n == 0 && m == 0) {
		if n == 0 {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}

		i--
	}
}

/*
根据上面的方法略微改进
	因为已知nums1的最后结果是装进两个已有切片，所以应该是当nums2遍历结束后num1就不需要遍历了，所以终止条件为nums2遍历结束，就不需要两次遍历判断是否都遍历结束
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i := len(nums1)
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i-1] = nums1[m-1]
			m--
		} else {
			nums1[i-1] = nums2[n-1]
			n--
		}
		i--
	}
}

func Test_T088(t *testing.T) {
	ints := []int{2, 0}
	merge(ints, 1, []int{1}, 1)
	fmt.Println(ints)
}
