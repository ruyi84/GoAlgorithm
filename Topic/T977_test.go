package main

// 直观法，先计算平方，再排序
func sortedSquares1(nums []int) []int {
	for i, num := range nums {
		nums[i] *= num
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

// 双指针法，
func sortedSquares(nums []int) []int {
	var mid int
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		mid = i
	}

	var newNums []int
	for left, right := mid, mid+1; left >= 0 || right < len(nums); {
		if left < 0 {
			newNums = append(newNums, nums[right]*nums[right])
			right++
			continue
		}
		if right >= len(nums) {
			newNums = append(newNums, nums[left]*nums[left])
			left--
			continue
		}
		if nums[left]*nums[left] > nums[right]*nums[right] {
			newNums = append(newNums, nums[right]*nums[right])
			right++
		} else {
			newNums = append(newNums, nums[left]*nums[left])
			left--
		}
	}

	return newNums
}
