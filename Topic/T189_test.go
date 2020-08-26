package main

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
*/

/*
第一次反转，直接将所有元素反转
然后两次反转，分别反转前k个元素，再反转后面len-k个元素
*/
func rotate(nums []int, k int) {
	reverse1(nums)
	reverse1(nums[:k%len(nums)])
	reverse1(nums[k%len(nums):])
}

func reverse1(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
