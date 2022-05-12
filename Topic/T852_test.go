package main

import (
	"fmt"
	"testing"
)

// 直接法，效率已经很高了
func peakIndexInMountainArray1(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return i - 1
		}
	}

	return len(arr) - 1
}

func TestT852(t *testing.T){
	//t.Log(peakIndexInMountainArray([]int{0,1,0}))
	t.Log(peakIndexInMountainArray([]int{3,5,3,2,0}))
	//t.Log(peakIndexInMountainArray([]int{18,29,38,59,98,100,99,98,90}))
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right+1) / 2
		fmt.Println(mid)
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid] < arr[mid-1]{
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return len(arr)-1
}