package main

import (
	"fmt"
	"testing"
)

func summaryRanges(nums []int) []string {
	var result []string

	//index := 0
	end := 0
	//for i := 0; i < len(nums)-1; i++ {
	//	//if i == len(nums)-1{
	//	//	number := fmt.Sprintf("%d", nums[index])
	//	//	result = append(result, number)
	//	//	continue
	//	//}
	//	end = i
	//	now := nums[i]
	//	next := nums[i+1]
	//
	//	if now+1 == next {
	//		continue
	//	}
	//
	//	if index == end {
	//		number := fmt.Sprintf("%d", nums[index])
	//		result = append(result, number)
	//	} else {
	//		number := fmt.Sprintf("%d->%d", nums[index], nums[end])
	//		result = append(result, number)
	//	}
	//
	//	index = i + 1
	//}

	for i := 0; i < len(nums)-1; i++ {
		index := i
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i]+1 == nums[j] {
				continue
			}
			if index == end {
				number := fmt.Sprintf("%d", nums[index])
				result = append(result, number)
			} else {
				number := fmt.Sprintf("%d->%d", nums[index], nums[end])
				result = append(result, number)
			}
			i = j
		}
	}
	return result
}

func Test228(t *testing.T) {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	fmt.Println(summaryRanges([]int{}))
	fmt.Println(summaryRanges([]int{0}))
}
