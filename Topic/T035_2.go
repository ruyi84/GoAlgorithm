package main

import "fmt"

func searchInsert1(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if mid == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right - 1
}

func main() {

	arr := []int{1, 3, 5, 6}
	fmt.Println(searchInsert1(arr, 7))
}

/*
这个题目较为简单，使用了常用的顺序比较和二分法，实际性能差距较小

顺序比较法，直接遍历数组，如果匹配到则直接返回当前下标
	如果未匹配，则将目标值和当前value比较，如果大于当前value值，则目标值的插入位置应该为当前下标。

二分法，常规使用方式
	定义左右边界下标，当左右界限不相等时，则取中间值，和目标值进行比较
	如果若目标值和中间值相等，则返回中间值下标
	若目标值与中间值不相等，则根据情况移动左右边界，如目标值大于中间值，则左边界修改为中间值+1，目标值小于中间值则相反

利用leetCode里：
解决方法	执行耗时	内存消耗
顺序比较	12ms		3.2mb
二分法		8ms			3.2mb

最坏结果时，二分法的速度的确是顺序比较的一倍。

*/
